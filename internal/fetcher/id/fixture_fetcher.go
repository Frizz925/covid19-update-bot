package id

import (
	"github.com/frizz925/covid19japan-chatbot/internal/country"
	"github.com/frizz925/covid19japan-chatbot/internal/data"
	idData "github.com/frizz925/covid19japan-chatbot/internal/data/id"
	"github.com/frizz925/covid19japan-chatbot/internal/fetcher"
)

const FIXTURE_FILE_UPDATE = "update.json"

type FixtureFetcher struct {
	fetcher.FixtureFetcher
}

func NewFixtureFetcher(dir string) *FixtureFetcher {
	return &FixtureFetcher{
		FixtureFetcher: fetcher.FixtureFetcher{
			Directory: dir,
			CountryID: country.ID_INDONESIA,
		},
	}
}

func (f *FixtureFetcher) Update() (*idData.UpdateResponse, error) {
	rc, err := f.ReadFile(FIXTURE_FILE_UPDATE)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return idData.ParseUpdate(rc)
}

func (f *FixtureFetcher) DailySummary() (*data.DailySummary, error) {
	ur, err := f.Update()
	if err != nil {
		return nil, err
	}
	return ur.Normalize()
}
