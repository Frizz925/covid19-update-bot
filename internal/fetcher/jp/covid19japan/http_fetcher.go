package covid19japan

import (
	"io"
	"net/http"

	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/data/jp/covid19japan"
)

const API_URL = "https://data.covid19japan.com/summary/latest.json"

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

func (hf *HTTPFetcher) SummaryLatest() (*covid19japan.SummaryLatest, error) {
	rc, err := hf.fetch(API_URL)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return covid19japan.ParseSummaryLatest(rc)
}

func (hf *HTTPFetcher) DailySummary() (*data.DailySummary, error) {
	sl, err := hf.SummaryLatest()
	if err != nil {
		return nil, err
	}
	return toNormalizedDailySummary(sl)
}

func (hf *HTTPFetcher) fetch(url string) (io.ReadCloser, error) {
	resp, err := hf.client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
