package routines

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
	"github.com/frizz925/covid19-update-bot/internal/formatters"
	"github.com/frizz925/covid19-update-bot/internal/publisher"
	"github.com/frizz925/covid19-update-bot/internal/scraper"
	"github.com/frizz925/covid19-update-bot/internal/templates"
)

var (
	ErrNotImplemented = errors.New("not yet implemented")
	ErrUnknownScraper = errors.New("unknown scraper type")
)

type DailyUpdateRoutine struct {
	generator *templates.Generator
}

type DailyUpdateRunConfig struct {
	Publisher publisher.Publisher
	Scraper   scraper.Scraper
}

func NewDailyUpdateRoutine() *DailyUpdateRoutine {
	return &DailyUpdateRoutine{}
}

func (r *DailyUpdateRoutine) Start(ctx context.Context, cfg *DailyUpdateRunConfig) error {
	switch v := cfg.Scraper.(type) {
	case scraper.ParsedScraper:
		return r.updateParsed(ctx, cfg, v)
	case scraper.ImageScraper:
		return r.updateImage(ctx, cfg, v)
	}
	return ErrUnknownScraper
}

func (r *DailyUpdateRoutine) updateParsed(ctx context.Context, cfg *DailyUpdateRunConfig, ps scraper.ParsedScraper) error {
	ds, err := ps.DailySummary(ctx)
	if err != nil {
		return err
	}
	return cfg.Publisher.PublishEmbed(&publisher.Embed{
		Author: publisher.Author{
			Name: r.createAuthorName(ds.Metadata),
			URL:  ps.Source(),
		},
		Title:       r.createTitle(ds.Country),
		Description: r.createDescription(ds.Date),
		Fields: []publisher.Field{
			{
				Name:  "Confirmed",
				Value: fmt.Sprintf("%s (+%s)", formatters.IntToNumber(ds.ConfirmedCumulative), formatters.IntToNumber(ds.Confirmed)),
			},
			{
				Name:  "Recovered",
				Value: fmt.Sprintf("%s (+%s)", formatters.IntToNumber(ds.RecoveredCumulative), formatters.IntToNumber(ds.Recovered)),
			},
			{
				Name:  "Deceased",
				Value: fmt.Sprintf("%s (+%s)", formatters.IntToNumber(ds.DeceasedCumulative), formatters.IntToNumber(ds.Deceased)),
			},
		},
		URL:       ds.Source.URL,
		Timestamp: ds.UpdatedAt,
	})
}

func (r *DailyUpdateRoutine) updateImage(ctx context.Context, cfg *DailyUpdateRunConfig, ps scraper.ImageScraper) error {
	dsi, err := ps.DailySummaryImage(ctx)
	if err != nil {
		return err
	}
	return cfg.Publisher.PublishEmbed(&publisher.Embed{
		Author: publisher.Author{
			Name: r.createAuthorName(dsi.Metadata),
			URL:  ps.Source(),
		},
		Title:       r.createTitle(dsi.Country),
		Description: r.createDescription(dsi.Date),
		ImageURL:    dsi.ImageURL,
		URL:         dsi.Source.URL,
		Timestamp:   dsi.UpdatedAt,
	})
}

func (DailyUpdateRoutine) createAuthorName(md data.Metadata) string {
	return fmt.Sprintf("%s %s", md.Country.Emoji(), md.Source.Comment)
}

func (DailyUpdateRoutine) createTitle(c country.Country) string {
	return fmt.Sprintf("Daily COVID-19 Update in %s", c.Name())
}

func (DailyUpdateRoutine) createDescription(date time.Time) string {
	return date.Format("Monday, January 2, 2006")
}
