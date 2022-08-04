package config

type Config struct {
	CountryID string  `json:"country_id"`
	Discord   Discord `json:"discord"`
}

type Discord struct {
	BotToken   string   `json:"bot_token"`
	ChannelIDs []string `json:"channel_ids"`
}
