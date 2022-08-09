package covid19goid

import (
	"github.com/frizz925/covid19-update-bot/internal/data/id/covid19goid"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	Update() (*covid19goid.UpdateResponse, error)
}
