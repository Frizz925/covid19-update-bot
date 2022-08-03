package data

import (
	"encoding/json"
	"io"
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
