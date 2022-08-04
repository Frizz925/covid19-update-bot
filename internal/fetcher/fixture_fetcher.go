package fetcher

import (
	"io"
	"os"
	"path"
)

type FixtureFetcher struct {
	Directory string
	CountryID string
}

func (f *FixtureFetcher) ReadFile(name string) (io.ReadCloser, error) {
	return os.Open(path.Join(f.Directory, f.CountryID, name))
}
