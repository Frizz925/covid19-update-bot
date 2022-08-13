package config

import (
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/scraper"
)

type Config struct {
	DataSources []DataSource `json:"data_sources"`
	Discord     Discord      `json:"discord"`
}

type Discord struct {
	BotToken   string   `json:"bot_token"`
	ChannelIDs []string `json:"channel_ids"`
}

type DataSource struct {
	Country     country.Country
	ScraperType scraper.Type
	Source      string
}
