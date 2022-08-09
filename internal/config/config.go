package config

type Config struct {
	DataSources map[string]string `json:"data_sources"`
	Discord     Discord           `json:"discord"`
}

type Discord struct {
	BotToken   string   `json:"bot_token"`
	ChannelIDs []string `json:"channel_ids"`
}
