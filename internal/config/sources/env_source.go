package sources

import (
	"context"
	"os"
	"strings"

	"github.com/frizz925/covid19-update-bot/internal/config"
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/scraper"
	"github.com/frizz925/covid19-update-bot/internal/storage"
)

const (
	ENV_COVID19_DATA_SOURCES = "COVID19_DATA_SOURCES"

	ENV_DISCORD_BOT_TOKEN   = "DISCORD_BOT_TOKEN"
	ENV_DISCORD_CHANNEL_IDS = "DISCORD_CHANNEL_IDS"

	ENV_STORAGE_TYPE      = "STORAGE_TYPE"
	ENV_STORAGE_S3_REGION = "STORAGE_S3_REGION"
	ENV_STORAGE_S3_BUCKET = "STORAGE_S3_BUCKET"
)

type envSource struct{}

func EnvSource() config.Source {
	return &envSource{}
}

func (es *envSource) Load(context.Context) (*config.Config, error) {
	return &config.Config{
		DataSources: es.getDataSources(),
		Discord: config.Discord{
			BotToken:   os.Getenv(ENV_DISCORD_BOT_TOKEN),
			ChannelIDs: strings.Split(os.Getenv(ENV_DISCORD_CHANNEL_IDS), ","),
		},
		Storage: config.Storage{
			Type:     storage.Type(os.Getenv(ENV_STORAGE_TYPE)),
			S3Region: os.Getenv(ENV_STORAGE_S3_REGION),
			S3Bucket: os.Getenv(ENV_STORAGE_S3_BUCKET),
		},
	}, nil
}

func (envSource) getDataSources() []config.DataSource {
	text := strings.TrimSpace(os.Getenv(ENV_COVID19_DATA_SOURCES))
	if text == "" {
		return nil
	}
	res := make([]config.DataSource, 0)
	for _, token := range strings.Split(text, ",") {
		toks := strings.SplitN(token, ":", 3)
		cid, st, src := toks[0], scraper.Parsed, ""
		if len(toks) >= 2 {
			st = scraper.Type(toks[1])
		}
		if len(toks) >= 3 {
			src = toks[2]
		}
		res = append(res, config.DataSource{
			Country:     country.Country(cid),
			ScraperType: st,
			Source:      src,
		})
	}
	return res
}
