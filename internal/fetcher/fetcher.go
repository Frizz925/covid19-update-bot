package fetcher

import (
	"errors"

	"github.com/frizz925/covid19-update-bot/internal/data"
)

type Type int

const (
	FixtureType Type = iota
	HTTPType
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not yet implemented")
)

type Fetcher interface {
	Type() Type
	Source() string
}

type ParsedFetcher interface {
	Fetcher
	DailySummary() (*data.DailySummary, error)
}

type ImageFetcher interface {
	Fetcher
	Image(url string) ([]byte, error)
}
