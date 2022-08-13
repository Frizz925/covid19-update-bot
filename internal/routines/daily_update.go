package routines

import (
	"context"
	"errors"
	"fmt"
	"net/url"

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

func DailyUpdate(ctx context.Context, cfg *DailyUpdateConfig) error {
	gen, err := templates.NewGenerator(cfg.TemplatesDir)
	if err != nil {
		return err
	}
	switch v := cfg.Scraper.(type) {
	case scraper.ParsedScraper:
		return dailyUpdateParsed(ctx, cfg, gen, v)
	case scraper.ImageScraper:
		return dailyUpdateImage(ctx, cfg, v)
	}
	return ErrUnknownScraper
}

func dailyUpdateParsed(ctx context.Context, cfg *DailyUpdateConfig, gen *templates.Generator, ps scraper.ParsedScraper) error {
	ds, err := ps.DailySummary(ctx)
	if err != nil {
		return err
	}
	message, err := gen.Daily(ds, fmt.Sprintf("Source: %s", ds.Source))
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}

func dailyUpdateImage(ctx context.Context, cfg *DailyUpdateConfig, ps scraper.ImageScraper) error {
	dsi, err := ps.DailySummaryImage(ctx)
	if err != nil {
		return err
	}
	u, err := url.Parse(ps.Source())
	if err != nil {
		return err
	}
	return cfg.Publisher.PublishEmbed(&publisher.Embed{
		Title:    fmt.Sprintf("Daily COVID-19 Update in %s", dsi.Country.Name()),
		Content:  dsi.DateTime.Format("Monday, January 2, 2006"),
		ImageURL: dsi.ImageURL,
		URL:      dsi.Source,
		Footer:   fmt.Sprintf("Data from %s • Updated at %s", u.Host, dsi.DateTime.Format("15:04 MST")),
	})
}
