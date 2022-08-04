package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/frizz925/covid19japan-chatbot/internal/config"
	"github.com/frizz925/covid19japan-chatbot/internal/data"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher/japan"
	"github.com/frizz925/covid19japan-chatbot/internal/publisher"
	"github.com/frizz925/covid19japan-chatbot/internal/routines"
	"github.com/joho/godotenv"
)

const (
	DIR_TEMPLATES = "templates"
	DIR_FIXTURES  = "fixtures"

	ENV_CHECK_AWS_LAMBDA = "LAMBDA_TASK_ROOT"
	ENV_CHECK_STAGING    = "DISCORD_BOT_STAGING"
)

type runConfig struct {
	lambdaEvent      *data.LambdaEvent
	fetchFromSource  bool
	publishToDiscord bool
}

func main() {
	if _, ok := os.LookupEnv(ENV_CHECK_AWS_LAMBDA); ok {
		lambda.Start(lambdaHandler)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	rcfg := &runConfig{}
	if os.Getenv(ENV_CHECK_STAGING) == "true" {
		rcfg.fetchFromSource = true
		rcfg.publishToDiscord = true
	}
	if err := run(ctx, rcfg); err != nil {
		panic(err)
	}
}

func lambdaHandler(ctx context.Context, event data.LambdaEvent) error {
	return run(ctx, &runConfig{
		lambdaEvent:      &event,
		fetchFromSource:  true,
		publishToDiscord: true,
	})
}

func run(ctx context.Context, cfg *runConfig) error {
	var fet fetcher.Fetcher
	if cfg.fetchFromSource {
		fet = japan.NewHTTPFetcher()
	} else {
		fet = japan.NewFixtureFetcher(DIR_FIXTURES)
	}

	var pub publisher.Publisher
	if cfg.publishToDiscord {
		var src config.Source
		if cfg.lambdaEvent != nil {
			src = config.AWSLambdaSource(cfg.lambdaEvent)
		} else {
			src = config.EnvSource()
		}
		cfg, err := src.Load(ctx)
		if err != nil {
			return err
		}
		dp, err := publisher.NewDiscordPublisher(&cfg.Discord)
		if err != nil {
			return err
		}
		if err := dp.Open(); err != nil {
			return err
		}
		pub = dp
	} else {
		pub = publisher.NewWritePublisher(os.Stdout)
	}
	defer cleanup(pub)

	return routines.DailyUpdate(&routines.DailyUpdateConfig{
		Fetcher:     fet,
		Publisher:   pub,
		TemplateDir: DIR_TEMPLATES,
	})
}

func cleanup(pub publisher.Publisher) {
	if v, ok := pub.(*publisher.DiscordPublisher); ok {
		v.Close()
	}
}
