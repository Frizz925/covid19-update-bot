package config

import (
	"context"
	"os"
	"strings"
)

const (
	ENV_COVID19_DATA_SOURCES = "COVID19_DATA_SOURCES"
	ENV_DISCORD_BOT_TOKEN    = "DISCORD_BOT_TOKEN"
	ENV_DISCORD_CHANNEL_IDS  = "DISCORD_CHANNEL_IDS"
)

type envSource struct{}

func EnvSource() Source {
	return &envSource{}
}

func (es *envSource) Load(context.Context) (*Config, error) {
	return &Config{
		DataSources: es.getDataSources(),
		Discord: Discord{
			BotToken:   os.Getenv(ENV_DISCORD_BOT_TOKEN),
			ChannelIDs: strings.Split(os.Getenv(ENV_DISCORD_CHANNEL_IDS), ","),
		},
	}, nil
}

func (envSource) getDataSources() map[string]string {
	text := strings.TrimSpace(os.Getenv(ENV_COVID19_DATA_SOURCES))
	if text == "" {
		return nil
	}
	res := make(map[string]string)
	for _, token := range strings.Split(text, ",") {
		toks := strings.SplitN(token, ":", 2)
		cid, src := toks[0], ""
		if len(toks) >= 2 {
			src = toks[1]
		}
		res[cid] = src
	}
	return res
}
