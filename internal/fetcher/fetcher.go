package fetcher

import (
	"context"
	"errors"
	"image"
	"net/url"

	"github.com/frizz925/covid19-update-bot/internal/data"
)

type Type int

const (
	FixtureType Type = iota
	HTTPType
)

var (
	ErrNotFound           = errors.New("not found")
	ErrNotImplemented     = errors.New("not yet implemented")
	ErrInvalidImageFormat = errors.New("invalid image format")
)

type Fetcher interface {
	Type() Type
	Source() string
	GetFullURL(url string) (*url.URL, error)
}

type ParsedFetcher interface {
	Fetcher
	DailySummary(ctx context.Context) (*data.DailySummary, error)
}

type ImageFetcher interface {
	Fetcher
	Image(ctx context.Context, url string) (image.Image, error)
}
