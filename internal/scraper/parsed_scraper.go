package scraper

import "github.com/frizz925/covid19-update-bot/internal/fetcher"

type parsedScraper struct {
	fetcher.ParsedFetcher
}

func NewParsedScraper(pf fetcher.ParsedFetcher) ParsedScraper {
	return &parsedScraper{pf}
}
