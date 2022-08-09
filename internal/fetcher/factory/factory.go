package factory

import (
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	idFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/id"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/id/covid19goid"
	jpFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/jp"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/jp/covid19japan"
)

type FetcherFactory struct {
	FixtureDir string
}

func NewFetcherFactory(fixtureDir string) *FetcherFactory {
	return &FetcherFactory{fixtureDir}
}

func (f *FetcherFactory) Fixture(countryID, source string) (fetcher.Fetcher, error) {
	switch countryID {
	case country.ID_JAPAN:
		switch source {
		case jpFetcher.DATA_SOURCE_COVID19JAPAN:
			return covid19japan.NewFixtureFetcher(f.FixtureDir), nil
		}
		return nil, fetcher.ErrNotImplemented
	case country.ID_INDONESIA:
		switch source {
		case idFetcher.DATA_SOURCE_COVID19_GO_ID:
			return covid19goid.NewFixtureFetcher(f.FixtureDir), nil
		}
		return nil, fetcher.ErrNotImplemented
	}
	return nil, fetcher.ErrNotFound
}

func (f *FetcherFactory) HTTP(countryID, source string) (fetcher.Fetcher, error) {
	switch countryID {
	case country.ID_JAPAN:
		switch source {
		case jpFetcher.DATA_SOURCE_COVID19JAPAN:
			return covid19japan.NewHTTPFetcher(), nil
		}
		return nil, fetcher.ErrNotImplemented
	case country.ID_INDONESIA:
		switch source {
		case idFetcher.DATA_SOURCE_COVID19_GO_ID:
			return covid19goid.NewHTTPFetcher(), nil
		}
		return nil, fetcher.ErrNotImplemented
	}
	return nil, fetcher.ErrNotFound
}
