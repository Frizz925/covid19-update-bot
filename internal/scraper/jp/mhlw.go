package jp

import (
	"encoding/xml"
	"errors"

	"github.com/frizz925/covid19-update-bot/internal/fetcher/jp/mhlw"
)

var ErrNotImplemented = errors.New("not yet implemented")

type MHLWScraper struct {
	mhlw.Fetcher
}

func NewMHLWScraper(mf mhlw.Fetcher) *MHLWScraper {
	return &MHLWScraper{mf}
}

func (ms *MHLWScraper) DailySummaryImage() ([]byte, error) {
	articleURL, err := ms.findArticleURL()
	if err != nil {
		return nil, err
	}
	_, err = ms.findSummaryImage(articleURL)
	if err != nil {
		return nil, err
	}
	return nil, ErrNotImplemented
}

func (ms *MHLWScraper) findArticleURL() (string, error) {
	rc, err := ms.Feed()
	if err != nil {
		return "", err
	}
	defer rc.Close()
	dec := xml.NewDecoder(rc)
	_, err = dec.Token()
	if err != nil {
		return "", err
	}
	return "", ErrNotImplemented
}

func (ms *MHLWScraper) findSummaryImage(url string) (string, error) {
	_, err := ms.News(url)
	if err != nil {
		return "", err
	}
	return "", ErrNotImplemented
}
