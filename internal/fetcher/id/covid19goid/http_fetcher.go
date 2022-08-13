package covid19goid

import (
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/id/covid19goid"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

const API_URL = "https://data.covid19.go.id/public/api/update.json"

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
	return API_URL
}

func (hf *HTTPFetcher) Update() (*covid19goid.UpdateResponse, error) {
	rc, err := hf.Fetch(hf.Source())
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
