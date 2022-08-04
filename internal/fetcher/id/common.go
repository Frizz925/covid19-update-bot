package id

import (
	idData "github.com/frizz925/covid19-update-bot/internal/data/id"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	Update() (*idData.UpdateResponse, error)
}
