package config

import (
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/scraper"
)

type Config struct {
	DataSources []DataSource `json:"data_sources"`
	Discord     Discord      `json:"discord"`
	Storage     Storage      `json:"storage"`
}

type DataSource struct {
	Country     country.Country `json:"country"`
	ScraperType scraper.Type    `json:"scraper_type"`
	Source      string          `json:"source"`
}

type Discord struct {
	BotToken   string   `json:"bot_token"`
	ChannelIDs []string `json:"channel_ids"`
}

type Storage struct {
	S3Region string `json:"s3_region"`
	S3Bucket string `json:"s3_bucket"`
}
