package jp

import (
	"encoding/xml"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
	mhlwData "github.com/frizz925/covid19-update-bot/internal/data/jp/mhlw"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/jp/mhlw"
)

var (
	ErrNotImplemented = errors.New("not yet implemented")
	ErrNotFound       = errors.New("not found")
)

type MHLWScraper struct {
	mhlw.Fetcher
}

type mhlwArticle struct {
	URL   string
	Title string
	Date  time.Time
}

func NewMHLWScraper(mf mhlw.Fetcher) *MHLWScraper {
	return &MHLWScraper{mf}
}

func (ms *MHLWScraper) DailySummaryImage() (*data.DailySummaryImage, error) {
	article, err := ms.findArticle()
	if err != nil {
		return nil, err
	}
	imgURL, err := ms.findSummaryImage(article.URL)
	if err != nil {
		return nil, err
	}
	img, err := ms.Image(imgURL)
	if err != nil {
		return nil, err
	}

	base := image.NewRGBA(img.Bounds())
	draw.Draw(base, base.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	draw.Draw(base, base.Bounds(), img, image.Point{}, draw.Over)

	return &data.DailySummaryImage{
		Country:  country.JP,
		DateTime: article.Date,
		Image:    base,
		Source:   article.URL,
	}, nil
}

func (ms *MHLWScraper) findArticle() (*mhlwArticle, error) {
	rc, err := ms.Feed()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	var rdf mhlwData.RDF
	if err := xml.NewDecoder(rc).Decode(&rdf); err != nil {
		return nil, err
	}
	var article mhlwArticle
	for _, item := range rdf.Items {
		if !strings.Contains(item.Title, "コロナウイルス感染症の現在") {
			continue
		}
		date, err := time.Parse("2006-01-02T15:04:05-07:00", item.Date)
		if err != nil {
			return nil, err
		}
		if !article.Date.IsZero() && article.Date.After(date) {
			continue
		}
		article.Title = item.Title
		article.Date = date
		article.URL = item.Link
	}
	if article.URL != "" {
		return &article, nil
	}
	return nil, ErrNotFound
}

func (ms *MHLWScraper) findSummaryImage(url string) (string, error) {
	rc, err := ms.News(url)
	if err != nil {
		return "", err
	}
	defer rc.Close()
	doc, err := goquery.NewDocumentFromReader(rc)
	if err != nil {
		return "", err
	}
	if src, ok := doc.Find(".l-contentBody img").First().Attr("src"); ok {
		return src, nil
	}
	return "", ErrNotFound
}
