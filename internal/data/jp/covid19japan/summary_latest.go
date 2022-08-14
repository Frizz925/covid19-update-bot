package covid19japan

import (
	"encoding/json"
	"io"
	"time"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
)

type SummaryLatest struct {
	Daily   []DailySummary `json:"daily"`
	Updated string         `json:"updated"`
	Source  data.Source    `json:"-"`
}

type DailySummary struct {
	Confirmed           int         `json:"confirmed"`
	Recovered           int         `json:"recovered"`
	Deceased            int         `json:"deceased"`
	ConfirmedCumulative int         `json:"confirmedCumulative"`
	RecoveredCumulative int         `json:"recoveredCumulative"`
	DeceasedCumulative  int         `json:"deceasedCumulative"`
	Date                string      `json:"date"`
	Updated             string      `json:"-"`
	Source              data.Source `json:"-"`
}

func ParseSummaryLatest(r io.Reader, source data.Source) (*SummaryLatest, error) {
	sl := SummaryLatest{Source: source}
	if err := sl.Parse(r); err != nil {
		return nil, err
	}
	return &sl, nil
}

func (sl *SummaryLatest) Parse(r io.Reader) error {
	return json.NewDecoder(r).Decode(sl)
}

func (sl *SummaryLatest) Today() *DailySummary {
	// Always assume the last daily summary is today's
	count := len(sl.Daily)
	if count <= 0 {
		return nil
	}
	ds := &sl.Daily[count-1]
	ds.Updated = sl.Updated
	ds.Source = sl.Source
	return ds
}

func (ds *DailySummary) Normalize() (*data.DailySummary, error) {
	date, err := parseDate(ds.Date)
	if err != nil {
		return nil, err
	}
	updated, err := parseUpdatedTime(ds.Updated)
	if err != nil {
		return nil, err
	}
	return &data.DailySummary{
		Metadata: data.Metadata{
			Country:   country.JP,
			Date:      date,
			UpdatedAt: updated,
			Source:    ds.Source,
		},
		Confirmed:           ds.Confirmed,
		Recovered:           ds.Recovered,
		Deceased:            ds.Deceased,
		ConfirmedCumulative: ds.ConfirmedCumulative,
		RecoveredCumulative: ds.RecoveredCumulative,
		DeceasedCumulative:  ds.DeceasedCumulative,
	}, nil
}

func parseDate(text string) (time.Time, error) {
	return time.Parse("2006-01-02", text)
}

func parseUpdatedTime(text string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", text)
}
