package fetcher

import (
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/frizz925/covid19-update-bot/internal/country"
)

const SOURCE_COMMENT_FIXTURE_SUFFIX = "(fixture)"

type FixtureFetcher struct {
	Directory  string
	Country    country.Country
	SourceName string
}

func (FixtureFetcher) Type() Type {
	return FixtureType
}

func (f *FixtureFetcher) Source() string {
	return f.GetPath("")
}

func (f *FixtureFetcher) ReadFile(name string) (io.ReadCloser, error) {
	return os.Open(f.GetPath(name))
}

func (f *FixtureFetcher) GetPath(name string) string {
	if path.IsAbs(name) {
		return name
	}
	return path.Join(f.Directory, f.Country.ID(), f.SourceName, name)
}

func (f *FixtureFetcher) GetAbsPath(name string) (string, error) {
	if path.IsAbs(name) {
		return name, nil
	}
	return filepath.Abs(f.GetPath(name))
}

func (f *FixtureFetcher) GetFullURL(rawURL string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	name := path.Base(u.Path)
	abs, err := f.GetAbsPath(name)
	if err != nil {
		return nil, err
	}
	return &url.URL{
		Scheme: "file",
		Path:   abs,
	}, nil
}
