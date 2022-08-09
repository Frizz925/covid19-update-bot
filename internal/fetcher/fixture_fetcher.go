package fetcher

import (
	"io"
	"os"
	"path"
)

type FixtureFetcher struct {
	Directory string
	CountryID string
	Source    string
}

func (f *FixtureFetcher) ReadFile(name string) (io.ReadCloser, error) {
	return os.Open(f.GetPath(name))
}

func (f *FixtureFetcher) GetPath(name string) string {
	return path.Join(f.Directory, f.CountryID, f.Source, name)
}
