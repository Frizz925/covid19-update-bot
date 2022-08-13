package covid19japan

import (
	"context"

	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/jp/covid19japan"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	SummaryLatest(ctx context.Context) (*covid19japan.SummaryLatest, error)
}

func toNormalizedDailySummary(sl *covid19japan.SummaryLatest) (*data.DailySummary, error) {
	ds := sl.Today()
	if ds == nil {
		return nil, fetcher.ErrNotFound
	}
	return ds.Normalize()
}
