package jp

import (
	"io"
	"net/http"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
	jpData "github.com/frizz925/covid19japan-chatbot/internal/data/jp"
)

const COVID19_JP_URL = "https://data.covid19japan.com/summary/latest.json"

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

func (hf *HTTPFetcher) SummaryLatest() (*jpData.SummaryLatest, error) {
	rc, err := hf.fetch(COVID19_JP_URL)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return jpData.ParseSummaryLatest(rc)
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
