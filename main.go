package main

import (
	"os"

	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
	"github.com/frizz925/covid19japan-chatbot/internal/publisher"
	"github.com/frizz925/covid19japan-chatbot/internal/routines"
	"github.com/joho/godotenv"
)

const (
	ENV_DISCORD_BOT_TOKEN  = "DISCORD_BOT_TOKEN"
	ENV_DISCORD_CHANNEL_ID = "DISCORD_CHANNEL_ID"

	DIR_TEMPLATES = "templates"
	DIR_FIXTURES  = "fixtures"

	DATA_SOURCE_BASE_URL = "https://data.covid19japan.com/"

	FETCH_FROM_FIXTURE = false
	PUBLISH_TO_STDOUT  = false
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	var fet fetcher.Fetcher
	if FETCH_FROM_FIXTURE {
		fet = fetcher.NewFixtureFetcher(DIR_FIXTURES)
	} else {
		var err error
		fet, err = fetcher.NewHTTPFetcher(DATA_SOURCE_BASE_URL)
		if err != nil {
			return err
		}
	}

	var pub publisher.Publisher
	if PUBLISH_TO_STDOUT {
		pub = publisher.NewWritePublisher(os.Stdout)
	} else {
		token := os.Getenv(ENV_DISCORD_BOT_TOKEN)
		channelID := os.Getenv(ENV_DISCORD_CHANNEL_ID)
		dp, err := publisher.NewDiscordPublisher(token, channelID)
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
