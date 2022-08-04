package japan

import (
	"path"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
	"github.com/frizz925/covid19japan-chatbot/internal/data/japan"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
)

const FIXTURE_FILE_SUMMARY_LATEST = "summary_latest.json"

type FixtureFetcher struct {
	fetcher.FixtureFetcher
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{
		FixtureFetcher: fetcher.FixtureFetcher{
			Directory: path.Join(dir, "japan"),
		},
	}
}

func (f *FixtureFetcher) SummaryLatest() (*japan.SummaryLatest, error) {
	rc, err := f.ReadFile(FIXTURE_FILE_SUMMARY_LATEST)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return japan.ParseSummaryLatest(rc)
}

func (f *FixtureFetcher) DailySummary() (*data.DailySummary, error) {
	sl, err := f.SummaryLatest()
	if err != nil {
		return nil, err
	}
	return toNormalizedDailySummary(sl)
}
