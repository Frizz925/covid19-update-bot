package factory

import (
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	idFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/id"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/id/covid19goid"
	jpFetcher "github.com/frizz925/covid19-update-bot/internal/fetcher/jp"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/jp/covid19japan"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/jp/mhlw"
)

type FetcherFactory struct {
	FixtureDir string
}

func NewFetcherFactory(fixtureDir string) *FetcherFactory {
	return &FetcherFactory{fixtureDir}
}

func (f *FetcherFactory) ParsedFetcher(ft fetcher.Type, c country.Country, source string) (fetcher.ParsedFetcher, error) {
	switch ft {
	case fetcher.FixtureType:
		switch c {
		case country.JP:
			switch source {
			case jpFetcher.DATA_SOURCE_COVID19JAPAN:
				fallthrough
			case "":
				return covid19japan.NewFixtureFetcher(f.FixtureDir), nil
			}
		case country.ID:
			switch source {
			case idFetcher.DATA_SOURCE_COVID19_GO_ID:
				fallthrough
			case "":
				return covid19goid.NewFixtureFetcher(f.FixtureDir), nil
			}
		}
	case fetcher.HTTPType:
		switch c {
		case country.JP:
			switch source {
			case jpFetcher.DATA_SOURCE_COVID19JAPAN:
				fallthrough
			case "":
				return covid19japan.NewHTTPFetcher(), nil
			}
		case country.ID:
			switch source {
			case idFetcher.DATA_SOURCE_COVID19_GO_ID:
				fallthrough
			case "":
				return covid19goid.NewHTTPFetcher(), nil
			}
		}
	}
	return nil, fetcher.ErrNotFound
}

func (f *FetcherFactory) ImageFetcher(ft fetcher.Type, c country.Country, source string) (fetcher.ImageFetcher, error) {
	switch ft {
	case fetcher.FixtureType:
		switch c {
		case country.JP:
			switch source {
			case jpFetcher.DATA_SOURCE_MHLW:
				fallthrough
			case "":
				return mhlw.NewFixtureFetcher(f.FixtureDir), nil
			}
		case country.ID:
			switch source {
			case idFetcher.DATA_SOURCE_TWITTER:
				fallthrough
			case "":
				return nil, fetcher.ErrNotImplemented
			}
		}
	case fetcher.HTTPType:
		switch c {
		case country.JP:
			switch source {
			case jpFetcher.DATA_SOURCE_MHLW:
				fallthrough
			case "":
				return mhlw.NewHTTPFetcher(), nil
			}
		case country.ID:
			switch source {
			case idFetcher.DATA_SOURCE_TWITTER:
				fallthrough
			case "":
				return nil, fetcher.ErrNotImplemented
			}
		}
	}
	return nil, fetcher.ErrNotFound
}
