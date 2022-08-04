package data

type LambdaEvent struct {
	CountryID  string   `json:"country_id"`
	ChannelIDs []string `json:"channel_ids"`
}
