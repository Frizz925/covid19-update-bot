package covid19japan

import (
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/jp/covid19japan"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
)

const API_URL = "https://data.covid19japan.com/summary/latest.json"

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

func (hf *HTTPFetcher) SummaryLatest() (*covid19japan.SummaryLatest, error) {
	resp, err := hf.Fetch(hf.Source())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return covid19japan.ParseSummaryLatest(resp.Body, hf.Source())
}

func (hf *HTTPFetcher) DailySummary() (*data.DailySummary, error) {
	sl, err := hf.SummaryLatest()
	if err != nil {
		return nil, err
	}
	return toNormalizedDailySummary(sl)
}
