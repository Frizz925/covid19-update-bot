package routines

import (
	"errors"
	"fmt"

	"github.com/frizz925/covid19-update-bot/internal/country"
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
	gen, err := templates.NewGenerator(cfg.TemplatesDir)
	if err != nil {
		return err
	}
	switch v := cfg.Scraper.(type) {
	case scraper.ParsedScraper:
		return dailyUpdateParsed(cfg, gen, v)
	case scraper.ImageScraper:
		return dailyUpdateImage(cfg, v)
	}
	return ErrUnknownScraper
}

func dailyUpdateParsed(cfg *DailyUpdateConfig, gen *templates.Generator, ps scraper.ParsedScraper) error {
	ds, err := ps.DailySummary()
	if err != nil {
		return err
	}
	message, err := gen.Daily(ds, fmt.Sprintf("Source: %s", ds.Source))
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}

func dailyUpdateImage(cfg *DailyUpdateConfig, ps scraper.ImageScraper) error {
	dsi, err := ps.DailySummaryImage()
	if err != nil {
		return err
	}
	return cfg.Publisher.PublishEmbed(&publisher.Embed{
		Title:   fmt.Sprintf("Daily COVID-19 update in %s", dsi.Country.Name()),
		Content: dsi.DateTime.Format("January 2, 2006"),
		URL:     dsi.Source,
	})
}
