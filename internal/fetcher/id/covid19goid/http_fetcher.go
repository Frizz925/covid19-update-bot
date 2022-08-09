package covid19goid

import (
	"io"
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/id/covid19goid"
)

const API_URL = "https://data.covid19.go.id/public/api/update.json"

type HTTPFetcher struct {
	client *http.Client
}

func NewHTTPFetcher(client ...*http.Client) *HTTPFetcher {
	hc := http.DefaultClient
	if len(client) > 0 {
		hc = client[0]
	}
	return &HTTPFetcher{
		client: hc,
	}
}

func (hf *HTTPFetcher) Source() string {
	return API_URL
}

func (hf *HTTPFetcher) Update() (*covid19goid.UpdateResponse, error) {
	rc, err := hf.fetch(API_URL)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return covid19goid.ParseUpdate(rc)
}

func (hf *HTTPFetcher) DailySummary() (*data.DailySummary, error) {
	ur, err := hf.Update()
	if err != nil {
		return nil, err
	}
	return ur.Normalize()
}

func (hf *HTTPFetcher) fetch(url string) (io.ReadCloser, error) {
	resp, err := hf.client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
