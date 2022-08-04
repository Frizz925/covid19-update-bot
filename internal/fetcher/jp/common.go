package jp

import (
	"github.com/frizz925/covid19japan-chatbot/internal/data"
	jpData "github.com/frizz925/covid19japan-chatbot/internal/data/jp"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
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
