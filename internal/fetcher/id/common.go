package id

import (
	idData "github.com/frizz925/covid19japan-chatbot/internal/data/id"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	Update() (*idData.UpdateResponse, error)
}
