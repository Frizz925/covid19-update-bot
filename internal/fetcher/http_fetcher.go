package fetcher

import (
	"context"
	"net/http"
	"net/url"
)

type HTTPFetcher struct {
	RawBaseURL string
	BaseURL    *url.URL
	Client     *http.Client
}

func (HTTPFetcher) Type() Type {
	return HTTPType
}

func (hf *HTTPFetcher) Fetch(ctx context.Context, rawURL string) (*http.Response, error) {
	u, err := hf.GetFullURL(rawURL)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return hf.GetClient().Do(req)
}

func (hf *HTTPFetcher) GetClient() *http.Client {
	if hf.Client != nil {
		return hf.Client
	}
	return http.DefaultClient
}

func (hf *HTTPFetcher) GetBaseURL() (*url.URL, error) {
	if hf.BaseURL != nil {
		return hf.BaseURL, nil
	}
	u, err := url.Parse(hf.RawBaseURL)
	if err != nil {
		return nil, err
	}
	hf.BaseURL = u
	return u, nil
}

func (hf *HTTPFetcher) GetFullURL(rawURL string) (*url.URL, error) {
	baseURL, err := hf.GetBaseURL()
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = baseURL.Scheme
	}
	if u.Host == "" {
		u.Host = baseURL.Host
	}
	if u.User == nil {
		u.User = baseURL.User
	}
	return u, nil
}
