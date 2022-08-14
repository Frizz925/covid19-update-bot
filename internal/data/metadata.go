package data

import (
	"time"

	"github.com/frizz925/covid19-update-bot/internal/country"
)

type Metadata struct {
	Country   country.Country
	Date      time.Time
	UpdatedAt time.Time
	Source    Source
}
