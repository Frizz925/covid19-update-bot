package jp

import (
	"github.com/frizz925/covid19-update-bot/internal/data"
	jpData "github.com/frizz925/covid19-update-bot/internal/data/jp"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

const FIXTURE_FILE_SUMMARY_LATEST = "summary_latest.json"

type FixtureFetcher struct {
	fetcher.FixtureFetcher
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{
		FixtureFetcher: fetcher.FixtureFetcher{
			Directory: dir,
			CountryID: "jp",
		},
	}
}

func (f *FixtureFetcher) SummaryLatest() (*jpData.SummaryLatest, error) {
	rc, err := f.ReadFile(FIXTURE_FILE_SUMMARY_LATEST)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return jpData.ParseSummaryLatest(rc)
}

func (f *FixtureFetcher) DailySummary() (*data.DailySummary, error) {
	sl, err := f.SummaryLatest()
	if err != nil {
		return nil, err
	}
	return toNormalizedDailySummary(sl)
}
