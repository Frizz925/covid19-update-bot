package fetcher

import (
	"io"
	"os"
	"path"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
)

const FIXTURE_FILE_SUMMARY_LATEST = "summary_latest.json"

type FixtureFetcher struct {
	dir string
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{dir}
}

func (f *FixtureFetcher) SummaryLatest() (*data.SummaryLatest, error) {
	rc, err := f.readFile(FIXTURE_FILE_SUMMARY_LATEST)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return data.ParseSummaryLatest(rc)
}

func (f *FixtureFetcher) readFile(name string) (io.ReadCloser, error) {
	return os.Open(path.Join(f.dir, name))
}
