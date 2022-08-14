package lambda

import "github.com/frizz925/covid19-update-bot/internal/config"

type Event struct {
	DataSources []config.DataSource `json:"data_sources"`
	Channels    []Channel           `json:"channels"`
	Storage     config.Storage      `json:"storage"`
}

type Channel struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
}
