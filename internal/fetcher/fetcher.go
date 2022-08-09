package fetcher

import (
	"errors"

	"github.com/frizz925/covid19-update-bot/internal/data"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not yet implemented")
)

type Fetcher interface {
	Source() string
	DailySummary() (*data.DailySummary, error)
}
