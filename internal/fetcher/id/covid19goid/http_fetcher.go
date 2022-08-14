package covid19goid

import (
	"context"
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/id/covid19goid"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

const (
	SOURCE_URL = "https://covid19.go.id/"
	API_URL    = "https://data.covid19.go.id/public/api/update.json"
)

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
	return SOURCE_URL
}

func (hf *HTTPFetcher) Update(ctx context.Context) (*covid19goid.UpdateResponse, error) {
	resp, err := hf.Fetch(ctx, API_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return covid19goid.ParseUpdate(resp.Body, data.Source{
		URL:     SOURCE_URL,
		DataURL: API_URL,
		Comment: SOURCE_COMMENT,
	})
}

func (hf *HTTPFetcher) DailySummary(ctx context.Context) (*data.DailySummary, error) {
	ur, err := hf.Update(ctx)
	if err != nil {
		return nil, err
	}
	return ur.Normalize()
}
