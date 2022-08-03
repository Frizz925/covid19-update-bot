package main

import (
	"context"
	"os"
	"time"

	"github.com/frizz925/covid19japan-chatbot/internal/config"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
	"github.com/frizz925/covid19japan-chatbot/internal/publisher"
	"github.com/frizz925/covid19japan-chatbot/internal/routines"
	"github.com/joho/godotenv"
)

const (
	DIR_TEMPLATES = "templates"
	DIR_FIXTURES  = "fixtures"

	FETCH_FROM_FIXTURE = true
	PUBLISH_TO_STDOUT  = false
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		return err
	}

	var fet fetcher.Fetcher
	if FETCH_FROM_FIXTURE {
		fet = fetcher.NewFixtureFetcher(DIR_FIXTURES)
	} else {
		fet = fetcher.NewHTTPFetcher()
	}

	var pub publisher.Publisher
	if PUBLISH_TO_STDOUT {
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
