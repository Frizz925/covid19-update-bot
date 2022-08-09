package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/frizz925/covid19-update-bot/internal/config"
	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/factory"
	"github.com/frizz925/covid19-update-bot/internal/publisher"
	"github.com/frizz925/covid19-update-bot/internal/routines"
	"github.com/joho/godotenv"
)

const (
	DIR_TEMPLATES = "templates"
	DIR_FIXTURES  = "fixtures"

	ENV_CHECK_AWS_LAMBDA   = "LAMBDA_TASK_ROOT"
	ENV_FETCH_FROM_WEB     = "FETCH_FROM_WEB"
	ENV_PUBLISH_TO_DISCORD = "PUBLISH_TO_DISCORD"
)

type runConfig struct {
	lambdaEvent      *data.LambdaEvent
	fetchFromWeb     bool
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
	if os.Getenv(ENV_FETCH_FROM_WEB) == "true" {
		rcfg.fetchFromWeb = true
	}
	if os.Getenv(ENV_PUBLISH_TO_DISCORD) == "true" {
		rcfg.publishToDiscord = true
	}
	if err := run(ctx, rcfg); err != nil {
		panic(err)
	}
}

func lambdaHandler(ctx context.Context, event data.LambdaEvent) error {
	return run(ctx, &runConfig{
		lambdaEvent:      &event,
		fetchFromWeb:     true,
		publishToDiscord: true,
	})
}

func run(ctx context.Context, rcfg *runConfig) error {
	var src config.Source
	if rcfg.lambdaEvent != nil {
		src = config.AWSLambdaSource(rcfg.lambdaEvent)
	} else {
		src = config.EnvSource()
	}
	cfg, err := src.Load(ctx)
	if err != nil {
		return err
	}

	var pub publisher.Publisher
	if rcfg.publishToDiscord {
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

	fac := factory.NewFetcherFactory(DIR_FIXTURES)
	for cid, src := range cfg.DataSources {
		var fet fetcher.Fetcher
		if rcfg.fetchFromWeb {
			fet, err = fac.HTTP(cid, src)
		} else {
			fet, err = fac.Fixture(cid, src)
		}
		if err != nil {
			return err
		}

		routineCfg := routines.DailyUpdateConfig{
			CountryID:   cid,
			Fetcher:     fet,
			Publisher:   pub,
			TemplateDir: DIR_TEMPLATES,
		}
		if err := routines.DailyUpdate(&routineCfg); err != nil {
			return err
		}
	}
	return nil
}

func cleanup(pub publisher.Publisher) {
	if v, ok := pub.(*publisher.DiscordPublisher); ok {
		v.Close()
	}
}
