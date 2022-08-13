package data

import (
	"time"

	"github.com/frizz925/covid19-update-bot/internal/country"
)

type DailySummary struct {
	Country             country.Country
	DateTime            time.Time
	Confirmed           int
	Recovered           int
	Deceased            int
	ConfirmedCumulative int
	RecoveredCumulative int
	DeceasedCumulative  int
}
