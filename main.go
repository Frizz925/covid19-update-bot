package main

import (
	"context"
	"os"
	"time"

	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/frizz925/covid19-update-bot/internal/config"
	"github.com/frizz925/covid19-update-bot/internal/config/sources"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	"github.com/frizz925/covid19-update-bot/internal/lambda"
	"github.com/frizz925/covid19-update-bot/internal/publisher"
	"github.com/frizz925/covid19-update-bot/internal/routines"
	"github.com/frizz925/covid19-update-bot/internal/scraper/factory"
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
	lambdaEvent      *lambda.Event
	fetcherType      fetcher.Type
	publishToDiscord bool
}

func main() {
	if _, ok := os.LookupEnv(ENV_CHECK_AWS_LAMBDA); ok {
		awsLambda.Start(lambdaHandler)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	rcfg := &runConfig{}
	if os.Getenv(ENV_FETCH_FROM_WEB) == "true" {
		rcfg.fetcherType = fetcher.HTTPType
	} else {
		rcfg.fetcherType = fetcher.FixtureType
	}
	if os.Getenv(ENV_PUBLISH_TO_DISCORD) == "true" {
		rcfg.publishToDiscord = true
	}
	if err := run(ctx, rcfg); err != nil {
		panic(err)
	}
}

func lambdaHandler(ctx context.Context, event lambda.Event) error {
	return run(ctx, &runConfig{
		lambdaEvent:      &event,
		fetcherType:      fetcher.HTTPType,
		publishToDiscord: true,
	})
}

func run(ctx context.Context, rcfg *runConfig) error {
	var src config.Source
	if rcfg.lambdaEvent != nil {
		src = sources.AWSLambdaSource(rcfg.lambdaEvent)
	} else {
		src = sources.EnvSource()
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

	fac := factory.NewScraperFactory(DIR_FIXTURES)
	for _, ds := range cfg.DataSources {
		scr, err := fac.Create(ds.ScraperType, rcfg.fetcherType, ds.Country, ds.Source)
		if err != nil {
			return err
		}

		routineCfg := routines.DailyUpdateConfig{
			TemplatesDir: DIR_TEMPLATES,
			Country:      ds.Country,
			Scraper:      scr,
			Publisher:    pub,
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
