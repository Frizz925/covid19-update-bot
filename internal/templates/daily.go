package templates

import (
	"time"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
	"github.com/frizz925/covid19japan-chatbot/internal/formatters"
)

const TEMPLATE_NAME_DAILY = "daily"

type dailyData struct {
	Confirmed           string
	Recovered           string
	Deceased            string
	ConfirmedCumulative string
	RecoveredCumulative string
	DeceasedCumulative  string
	Date                string
}

func (g *Generator) Daily(ds *data.DailySummary) (string, error) {
	date, err := parseDailyDate(ds.Date)
	if err != nil {
		return "", err
	}
	return g.Generate(TEMPLATE_NAME_DAILY, &dailyData{
		Confirmed:           formatters.IntToNumber(ds.Confirmed),
		Recovered:           formatters.IntToNumber(ds.Recovered),
		Deceased:            formatters.IntToNumber(ds.Deceased),
		ConfirmedCumulative: formatters.IntToNumber(ds.ConfirmedCumulative),
		RecoveredCumulative: formatters.IntToNumber(ds.RecoveredCumulative),
		DeceasedCumulative:  formatters.IntToNumber(ds.DeceasedCumulative),
		Date:                date.Format("Monday, January 2, 2006"),
	})
}

func parseDailyDate(text string) (time.Time, error) {
	return time.Parse("2006-01-02", text)
}
