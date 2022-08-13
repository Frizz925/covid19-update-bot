package routines

import (
	"errors"
	"fmt"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	"github.com/frizz925/covid19-update-bot/internal/publisher"
	"github.com/frizz925/covid19-update-bot/internal/scraper"
	"github.com/frizz925/covid19-update-bot/internal/templates"
)

var (
	ErrNotImplemented = errors.New("not yet implemented")
	ErrUnknownScraper = errors.New("unknown scraper type")
)

type DailyUpdateConfig struct {
	TemplatesDir string
	Country      country.Country
	Publisher    publisher.Publisher
	Scraper      scraper.Scraper
}

func DailyUpdate(cfg *DailyUpdateConfig) error {
	switch v := cfg.Scraper.(type) {
	case fetcher.ParsedFetcher:
		return dailyUpdateParsed(cfg, v)
	case scraper.Scraper:
		return ErrNotImplemented
	}
	return ErrUnknownScraper
}

func dailyUpdateParsed(cfg *DailyUpdateConfig, pf fetcher.ParsedFetcher) error {
	gen, err := templates.NewGenerator(cfg.TemplatesDir)
	if err != nil {
		return err
	}
	ds, err := pf.DailySummary()
	if err != nil {
		return err
	}
	message, err := gen.Daily(ds, fmt.Sprintf("Source: %s", pf.Source()))
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}
