package covid19japan

import (
	"context"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/jp/covid19japan"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	jpFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/jp"
)

const FIXTURE_FILE_SUMMARY_LATEST = "summary_latest.json"

type FixtureFetcher struct {
	fetcher.FixtureFetcher
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{
		FixtureFetcher: fetcher.FixtureFetcher{
			Directory:  dir,
			Country:    country.JP,
			SourceName: jpFetcher.DATA_SOURCE_COVID19JAPAN,
		},
	}
}

func (f *FixtureFetcher) Source() string {
	return f.GetPath("")
}

func (f *FixtureFetcher) SummaryLatest(ctx context.Context) (*covid19japan.SummaryLatest, error) {
	u, err := f.GetFullURL(FIXTURE_FILE_SUMMARY_LATEST)
	if err != nil {
		return nil, err
	}
	rc, err := f.ReadFile(u.Path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return covid19japan.ParseSummaryLatest(rc, data.Source{
		URL:     SOURCE_URL,
		DataURL: u.String(),
		Comment: SOURCE_COMMENT + " " + fetcher.SOURCE_COMMENT_FIXTURE_SUFFIX,
	})
}

func (f *FixtureFetcher) DailySummary(ctx context.Context) (*data.DailySummary, error) {
	sl, err := f.SummaryLatest(ctx)
	if err != nil {
		return nil, err
	}
	return toNormalizedDailySummary(sl)
}
