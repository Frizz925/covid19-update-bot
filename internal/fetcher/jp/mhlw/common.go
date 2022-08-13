package mhlw

import (
	"io"

	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

type Fetcher interface {
	fetcher.Fetcher
	Feed() (io.ReadCloser, error)
	News(url string) (io.ReadCloser, error)
	Image(url string) ([]byte, error)
}
