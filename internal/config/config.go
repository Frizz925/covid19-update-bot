package config

type Config struct {
	Discord Discord `json:"discord"`
}

type Discord struct {
	BotToken   string   `json:"bot_token"`
	ChannelIDs []string `json:"channel_ids"`
}
