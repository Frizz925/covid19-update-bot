package mhlw

import (
	"io"
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

const FEED_URL = "https://www.mhlw.go.jp/stf/news.rdf"

type HTTPFetcher struct {
	fetcher.HTTPFetcher
}

func NewHTTPFetcher(client ...*http.Client) *HTTPFetcher {
	hf := &HTTPFetcher{}
	if len(client) > 0 {
		hf.Client = client[0]
	}
	return hf
}

func (hf *HTTPFetcher) Source() string {
	return FEED_URL
}

func (hf *HTTPFetcher) Feed() (io.ReadCloser, error) {
	return hf.Fetch(hf.Source())
}

func (hf *HTTPFetcher) News(url string) (io.ReadCloser, error) {
	return hf.Fetch(url)
}

func (hf *HTTPFetcher) Image(url string) ([]byte, error) {
	rc, err := hf.Fetch(url)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}
