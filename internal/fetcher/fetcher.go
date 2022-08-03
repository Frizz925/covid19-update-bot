package fetcher

import "github.com/frizz925/covid19japan-chatbot/internal/data"

type Fetcher interface {
	SummaryLatest() (*data.SummaryLatest, error)
}
