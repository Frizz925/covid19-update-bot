package covid19japan

import (
	"encoding/json"
	"io"
	"time"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
)

type SummaryLatest struct {
	Daily  []DailySummary `json:"daily"`
	Source string         `json:"-"`
}

type DailySummary struct {
	Confirmed           int    `json:"confirmed"`
	Recovered           int    `json:"recovered"`
	Deceased            int    `json:"deceased"`
	ConfirmedCumulative int    `json:"confirmedCumulative"`
	RecoveredCumulative int    `json:"recoveredCumulative"`
	DeceasedCumulative  int    `json:"deceasedCumulative"`
	Date                string `json:"date"`
	Source              string `json:"-"`
}

func ParseSummaryLatest(r io.Reader, source string) (*SummaryLatest, error) {
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
	ds.Source = sl.Source
	return ds
}

func (ds *DailySummary) Normalize() (*data.DailySummary, error) {
	date, err := parseDailyDate(ds.Date)
	if err != nil {
		return nil, err
	}
	return &data.DailySummary{
		Country:             country.JP,
		DateTime:            date,
		Confirmed:           ds.Confirmed,
		Recovered:           ds.Recovered,
		Deceased:            ds.Deceased,
		ConfirmedCumulative: ds.ConfirmedCumulative,
		RecoveredCumulative: ds.RecoveredCumulative,
		DeceasedCumulative:  ds.DeceasedCumulative,
		Source:              ds.Source,
	}, nil
}

func parseDailyDate(text string) (time.Time, error) {
	return time.Parse("2006-01-02", text)
}
