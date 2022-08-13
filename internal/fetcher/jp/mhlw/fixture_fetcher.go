package mhlw

import (
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/url"
	"path"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	jpFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/jp"
)

type FixtureFetcher struct {
	fetcher.FixtureFetcher
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{
		FixtureFetcher: fetcher.FixtureFetcher{
			Directory:  dir,
			Country:    country.JP,
			SourceName: jpFetcher.DATA_SOURCE_MHLW,
		},
	}
}

func (f *FixtureFetcher) Source() string {
	return f.GetPath("")
}

func (f *FixtureFetcher) Feed(ctx context.Context) (io.ReadCloser, error) {
	return f.ReadFile("news.rdf")
}

func (f *FixtureFetcher) News(ctx context.Context, rawURL string) (io.ReadCloser, error) {
	name, err := f.urlToFilename(rawURL)
	if err != nil {
		return nil, err
	}
	return f.ReadFile(name)
}

func (f *FixtureFetcher) Image(ctx context.Context, rawURL string) (image.Image, error) {
	name, err := f.urlToFilename(rawURL)
	if err != nil {
		return nil, err
	}
	rc, err := f.ReadFile(name)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	ext := path.Ext(name)
	switch ext {
	case ".png":
		return png.Decode(rc)
	case ".jpg":
		fallthrough
	case ".jpeg":
		return jpeg.Decode(rc)
	}
	return nil, fetcher.ErrInvalidImageFormat
}

func (f *FixtureFetcher) urlToFilename(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return path.Base(u.Path), nil
}
