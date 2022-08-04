package factory

import (
	"github.com/frizz925/covid19japan-chatbot/internal/country"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
	idFetcher "github.com/frizz925/covid19japan-chatbot/internal/fetcher/id"
	jpFetcher "github.com/frizz925/covid19japan-chatbot/internal/fetcher/jp"
)

type FetcherFactory struct {
	FixtureDir string
}

func NewFetcherFactory(fixtureDir string) *FetcherFactory {
	return &FetcherFactory{fixtureDir}
}

func (f *FetcherFactory) Fixture(countryID string) (fetcher.Fetcher, error) {
	switch countryID {
	case country.ID_JAPAN:
		return jpFetcher.NewFixtureFetcher(f.FixtureDir), nil
	case country.ID_INDONESIA:
		return idFetcher.NewFixtureFetcher(f.FixtureDir), nil
	}
	return nil, fetcher.ErrNotFound
}

func (f *FetcherFactory) HTTP(countryID string) (fetcher.Fetcher, error) {
	switch countryID {
	case country.ID_JAPAN:
		return jpFetcher.NewHTTPFetcher(), nil
	case country.ID_INDONESIA:
		return idFetcher.NewHTTPFetcher(), nil
	}
	return nil, fetcher.ErrNotFound
}
