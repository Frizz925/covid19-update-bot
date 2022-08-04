package templates

import (
	"strings"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
	"github.com/frizz925/covid19japan-chatbot/internal/formatters"
)

const TEMPLATE_NAME_DAILY = "daily"

type dailyData struct {
	Country             string
	Date                string
	Confirmed           string
	Recovered           string
	Deceased            string
	ConfirmedCumulative string
	RecoveredCumulative string
	DeceasedCumulative  string
	Comment             string
}

func (g *Generator) Daily(ds *data.DailySummary, comments ...string) (string, error) {
	res, err := g.Generate(TEMPLATE_NAME_DAILY, &dailyData{
		Country:             ds.Country,
		Date:                ds.DateTime.Format("Monday, January 2, 2006"),
		Confirmed:           formatters.IntToNumber(ds.Confirmed),
		Recovered:           formatters.IntToNumber(ds.Recovered),
		Deceased:            formatters.IntToNumber(ds.Deceased),
		ConfirmedCumulative: formatters.IntToNumber(ds.ConfirmedCumulative),
		RecoveredCumulative: formatters.IntToNumber(ds.RecoveredCumulative),
		DeceasedCumulative:  formatters.IntToNumber(ds.DeceasedCumulative),
		Comment:             strings.Join(comments, "\n"),
	})
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(res), nil
}
