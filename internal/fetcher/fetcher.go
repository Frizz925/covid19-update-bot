package fetcher

import (
	"errors"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
)

var ErrNotFound = errors.New("not found")

type Fetcher interface {
	DailySummary() (*data.DailySummary, error)
}
