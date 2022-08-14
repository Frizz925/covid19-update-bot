package data

type DailySummary struct {
	Metadata
	Confirmed           int
	Recovered           int
	Deceased            int
	ConfirmedCumulative int
	RecoveredCumulative int
	DeceasedCumulative  int
}

type DailySummaryImage struct {
	Metadata
	ImageURL string
}
