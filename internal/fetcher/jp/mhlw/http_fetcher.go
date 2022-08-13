package mhlw

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

const (
	BASE_URL  = "https://www.mhlw.go.jp"
	FEED_PATH = "/stf/news.rdf"
)

type HTTPFetcher struct {
	fetcher.HTTPFetcher
}

func NewHTTPFetcher(client ...*http.Client) *HTTPFetcher {
	hf := &HTTPFetcher{
		HTTPFetcher: fetcher.HTTPFetcher{
			RawBaseURL: BASE_URL,
		},
	}
	if len(client) > 0 {
		hf.Client = client[0]
	}
	return hf
}

func (hf *HTTPFetcher) Source() string {
	return BASE_URL
}

func (hf *HTTPFetcher) Feed() (io.ReadCloser, error) {
	resp, err := hf.Fetch(FEED_PATH)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (hf *HTTPFetcher) News(url string) (io.ReadCloser, error) {
	resp, err := hf.Fetch(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (hf *HTTPFetcher) Image(url string) (image.Image, error) {
	resp, err := hf.Fetch(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.Header.Get("Content-Type") {
	case "image/png":
		return png.Decode(resp.Body)
	case "image/jpg":
	case "image/jpeg":
		return jpeg.Decode(resp.Body)
	}
	return nil, fetcher.ErrInvalidImageFormat
}
