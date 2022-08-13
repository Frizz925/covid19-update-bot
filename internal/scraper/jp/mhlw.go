package jp

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"image/png"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
	mhlwData "github.com/frizz925/covid19-update-bot/internal/data/jp/mhlw"
	"github.com/frizz925/covid19-update-bot/internal/fetcher/jp/mhlw"
	"github.com/frizz925/covid19-update-bot/internal/imgproc"
	"github.com/frizz925/covid19-update-bot/internal/storage"
)

var (
	ErrNotImplemented = errors.New("not yet implemented")
	ErrNotFound       = errors.New("not found")
)

type MHLWScraper struct {
	mhlw.Fetcher
	storage.Storage
}

type mhlwArticle struct {
	URL   string
	Title string
	Date  time.Time
}

func NewMHLWScraper(mf mhlw.Fetcher, st storage.Storage) *MHLWScraper {
	return &MHLWScraper{
		Fetcher: mf,
		Storage: st,
	}
}

func (ms *MHLWScraper) DailySummaryImage(ctx context.Context) (*data.DailySummaryImage, error) {
	article, err := ms.findArticle(ctx)
	if err != nil {
		return nil, err
	}
	imgURL, err := ms.findSummaryImage(ctx, article.URL)
	if err != nil {
		return nil, err
	}
	img, err := ms.Image(ctx, imgURL)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, imgproc.WhiteBG(img)); err != nil {
		return nil, err
	}

	imgName := fmt.Sprintf("jp/mhlw/%s.png", article.Date.Format("2006-01-02"))
	obj, err := ms.Write(ctx, imgName, &buf)
	if err != nil {
		return nil, err
	}

	return &data.DailySummaryImage{
		Country:  country.JP,
		DateTime: article.Date,
		ImageURL: obj.URL(),
		Source:   article.URL,
	}, nil
}

func (ms *MHLWScraper) findArticle(ctx context.Context) (*mhlwArticle, error) {
	rc, err := ms.Feed(ctx)
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

func (ms *MHLWScraper) findSummaryImage(ctx context.Context, url string) (string, error) {
	rc, err := ms.News(ctx, url)
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
