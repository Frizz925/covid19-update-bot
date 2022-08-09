package routines

import (
	"fmt"

	"github.com/frizz925/covid19-update-bot/internal/fetcher"
	"github.com/frizz925/covid19-update-bot/internal/publisher"
	"github.com/frizz925/covid19-update-bot/internal/templates"
)

type DailyUpdateConfig struct {
	CountryID   string
	Source      string
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
	message, err := gen.Daily(ds, fmt.Sprintf("Source: %s", cfg.Fetcher.Source()))
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}
