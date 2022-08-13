package sources

import (
	"context"
	"os"
	"strings"

	"github.com/frizz925/covid19-update-bot/internal/config"
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/scraper"
)

const (
	ENV_COVID19_DATA_SOURCES = "COVID19_DATA_SOURCES"
	ENV_DISCORD_BOT_TOKEN    = "DISCORD_BOT_TOKEN"
	ENV_DISCORD_CHANNEL_IDS  = "DISCORD_CHANNEL_IDS"
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
