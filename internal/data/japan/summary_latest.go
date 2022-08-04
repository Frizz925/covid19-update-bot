package japan

import (
	"encoding/json"
	"io"
	"time"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
)

type SummaryLatest struct {
	Daily []DailySummary `json:"daily"`
}

type DailySummary struct {
	Confirmed           int    `json:"confirmed"`
	Recovered           int    `json:"recovered"`
	Deceased            int    `json:"deceased"`
	ConfirmedCumulative int    `json:"confirmedCumulative"`
	RecoveredCumulative int    `json:"recoveredCumulative"`
	DeceasedCumulative  int    `json:"deceasedCumulative"`
	Date                string `json:"date"`
}

func ParseSummaryLatest(r io.Reader) (*SummaryLatest, error) {
	var sl SummaryLatest
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
	return &sl.Daily[count-1]
}

func (ds *DailySummary) Normalize() (*data.DailySummary, error) {
	date, err := parseDailyDate(ds.Date)
	if err != nil {
		return nil, err
	}
	return &data.DailySummary{
		Country:             "Japan",
		CountryID:           "JP",
		DateTime:            date,
		Confirmed:           ds.Confirmed,
		Recovered:           ds.Recovered,
		Deceased:            ds.Deceased,
		ConfirmedCumulative: ds.ConfirmedCumulative,
		RecoveredCumulative: ds.RecoveredCumulative,
		DeceasedCumulative:  ds.DeceasedCumulative,
	}, nil
}

func parseDailyDate(text string) (time.Time, error) {
	return time.Parse("2006-01-02", text)
}
