package scraper

import "github.com/frizz925/covid19-update-bot/internal/data"

type Type string

const (
	Parsed Type = "txt"
	Image  Type = "img"
)

type Scraper interface {
	Source() string
}

type ParsedScraper interface {
	Scraper
	DailySummary() (*data.DailySummary, error)
}

type ImageScraper interface {
	Scraper
	DailySummaryImage() ([]byte, error)
}
