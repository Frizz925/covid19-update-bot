package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/frizz925/covid19japan-chatbot/internal/config"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
	"github.com/frizz925/covid19japan-chatbot/internal/publisher"
	"github.com/frizz925/covid19japan-chatbot/internal/routines"
	"github.com/joho/godotenv"
)

const (
	DIR_TEMPLATES = "templates"
	DIR_FIXTURES  = "fixtures"

	AWS_LAMBDA_ENV_CHECK = "LAMBDA_TASK_ROOT"
)

type runConfig struct {
	fetchFromFixture bool
	publishToStdout  bool
}

func main() {
	if _, ok := os.LookupEnv(AWS_LAMBDA_ENV_CHECK); ok {
		lambda.Start(lambdaHandler)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	rcfg := &runConfig{
		fetchFromFixture: true,
		publishToStdout:  true,
	}
	if err := run(ctx, rcfg); err != nil {
		panic(err)
	}
}

func lambdaHandler(ctx context.Context) error {
	return run(ctx, &runConfig{
		fetchFromFixture: false,
		publishToStdout:  false,
	})
}

func run(ctx context.Context, cfg *runConfig) error {
	var fet fetcher.Fetcher
	if cfg.fetchFromFixture {
		fet = fetcher.NewFixtureFetcher(DIR_FIXTURES)
	} else {
		fet = fetcher.NewHTTPFetcher()
	}

	var pub publisher.Publisher
	if cfg.publishToStdout {
		pub = publisher.NewWritePublisher(os.Stdout)
	} else {
		cfg, err := config.NewEnvSource().Load(ctx)
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
	}
	defer func(pub publisher.Publisher) {
		if v, ok := pub.(*publisher.DiscordPublisher); ok {
			v.Close()
		}
	}(pub)

	return routines.DailyUpdate(&routines.DailyUpdateConfig{
		Fetcher:     fet,
		Publisher:   pub,
		TemplateDir: DIR_TEMPLATES,
	})
}
