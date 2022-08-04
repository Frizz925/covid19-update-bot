package jp

import (
	"github.com/frizz925/covid19-update-bot/internal/data"
	jpData "github.com/frizz925/covid19-update-bot/internal/data/jp"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	SummaryLatest() (*jpData.SummaryLatest, error)
}

func toNormalizedDailySummary(sl *jpData.SummaryLatest) (*data.DailySummary, error) {
	ds := sl.Today()
	if ds == nil {
		return nil, fetcher.ErrNotFound
	}
	return ds.Normalize()
}
