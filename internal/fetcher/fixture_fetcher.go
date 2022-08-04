package fetcher

import (
	"io"
	"os"
	"path"
)

type FixtureFetcher struct {
	Directory string
}

func (f *FixtureFetcher) ReadFile(name string) (io.ReadCloser, error) {
	return os.Open(path.Join(f.Directory, name))
}
