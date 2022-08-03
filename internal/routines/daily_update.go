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
	sl, err := cfg.Fetcher.SummaryLatest()
	if err != nil {
		return err
	}
	gen, err := templates.NewGenerator(cfg.TemplateDir)
	if err != nil {
		return err
	}
	message, err := gen.Daily(sl.Today())
	if err != nil {
		return err
	}
	return cfg.Publisher.Publish(message)
}
