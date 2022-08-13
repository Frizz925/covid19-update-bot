package mhlw

import (
	"context"
	"image"
	"io"

	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	Feed(ctx context.Context) (io.ReadCloser, error)
	News(ctx context.Context, url string) (io.ReadCloser, error)
	Image(ctx context.Context, url string) (image.Image, error)
}
