package fetcher

import (
	"io"
	"net/http"
	"net/url"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
)

const (
	HTTP_PATH_SUMMARY_LATEST = "/summary/latest.json"
)

type HTTPFetcher struct {
	baseURL *url.URL
	client  *http.Client
}

func NewHTTPFetcher(baseURL string, client ...*http.Client) (*HTTPFetcher, error) {
	hc := http.DefaultClient
	if len(client) > 0 {
		hc = client[0]
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &HTTPFetcher{
		baseURL: normalizeURL(u),
		client:  hc,
	}, nil
}

func (hf *HTTPFetcher) SummaryLatest() (*data.SummaryLatest, error) {
	rc, err := hf.fetch(HTTP_PATH_SUMMARY_LATEST)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return data.ParseSummaryLatest(rc)
}

func (hf *HTTPFetcher) fetch(path string) (io.ReadCloser, error) {
	resp, err := hf.client.Get(createURL(hf.baseURL, path).String())
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func normalizeURL(u *url.URL) *url.URL {
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return u
}

func createURL(baseURL *url.URL, path string) *url.URL {
	return &url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   path,
	}
}
