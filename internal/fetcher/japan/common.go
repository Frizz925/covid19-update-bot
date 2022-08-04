package japan

import (
	"github.com/frizz925/covid19japan-chatbot/internal/data"
	"github.com/frizz925/covid19japan-chatbot/internal/data/japan"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
)

type JapanFetcher interface {
	fetcher.Fetcher
	SummaryLatest() (*japan.SummaryLatest, error)
}

func toNormalizedDailySummary(sl *japan.SummaryLatest) (*data.DailySummary, error) {
	ds := sl.Today()
	if ds == nil {
		return nil, fetcher.ErrNotFound
	}
	return ds.Normalize()
}
