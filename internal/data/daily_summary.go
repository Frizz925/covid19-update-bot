package data

import "time"

type DailySummary struct {
	Country             string
	CountryID           string
	DateTime            time.Time
	Confirmed           int
	Recovered           int
	Deceased            int
	ConfirmedCumulative int
	RecoveredCumulative int
	DeceasedCumulative  int
}
