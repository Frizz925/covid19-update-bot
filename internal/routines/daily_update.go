package routines

import (
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	"github.com/frizz925/covid19-update-bot/internal/publisher"
	"github.com/frizz925/covid19-update-bot/internal/templates"
)

type DailyUpdateConfig struct {
	CountryID   string
	TemplateDir string
	Fetcher     fetcher.Fetcher
	Publisher   publisher.Publisher
}

func DailyUpdate(cfg *DailyUpdateConfig) error {
	gen, err := templates.NewGenerator(cfg.TemplateDir)
	if err != nil {
		return err
	}
	ds, err := cfg.Fetcher.DailySummary()
	if err != nil {
		return err
	}
	message, err := gen.Daily(ds, country.Comments[cfg.CountryID])
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}
