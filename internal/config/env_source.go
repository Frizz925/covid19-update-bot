package config

import (
	"context"
	"os"
	"strings"
)

const (
	ENV_DISCORD_BOT_TOKEN   = "DISCORD_BOT_TOKEN"
	ENV_DISCORD_CHANNEL_IDS = "DISCORD_CHANNEL_IDS"
)

type envSource struct{}

func EnvSource() Source {
	return &envSource{}
}

func (*envSource) Load(context.Context) (*Config, error) {
	return &Config{
		Discord: Discord{
			BotToken:   os.Getenv(ENV_DISCORD_BOT_TOKEN),
			ChannelIDs: strings.Split(os.Getenv(ENV_DISCORD_CHANNEL_IDS), ","),
		},
	}, nil
}
