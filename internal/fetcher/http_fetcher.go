package fetcher

import (
	"io"
	"net/http"
)

type HTTPFetcher struct {
	Client *http.Client
}

func (HTTPFetcher) Type() Type {
	return HTTPType
}

func (hf *HTTPFetcher) Fetch(url string) (io.ReadCloser, error) {
	resp, err := hf.GetClient().Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (hf *HTTPFetcher) GetClient() *http.Client {
	if hf.Client != nil {
		return hf.Client
	}
	return http.DefaultClient
}
