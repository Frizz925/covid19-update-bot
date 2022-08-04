package routines

import (
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
	"github.com/frizz925/covid19japan-chatbot/internal/publisher"
	"github.com/frizz925/covid19japan-chatbot/internal/templates"
)

type DailyUpdateConfig struct {
	Fetcher     fetcher.Fetcher
	Publisher   publisher.Publisher
	TemplateDir string
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
	message, err := gen.Daily(ds, "Data from https://covid19japan.com/")
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}
