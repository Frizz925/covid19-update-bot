package covid19goid

import (
	"encoding/json"
	"io"
	"time"

	"github.com/frizz925/covid19-update-bot/internal/country"
	"github.com/frizz925/covid19-update-bot/internal/data"
)

type UpdateResponse struct {
	Update Update      `json:"update"`
	Source data.Source `json:"-"`
}

type Update struct {
	Penambahan Penambahan `json:"penambahan"`
	Total      Total      `json:"total"`
}

type Penambahan struct {
	JumlahPositif   int    `json:"jumlah_positif"`
	JumlahMeninggal int    `json:"jumlah_meninggal"`
	JumlahSembuh    int    `json:"jumlah_sembuh"`
	JumlahDirawat   int    `json:"jumlah_dirawat"`
	Tanggal         string `json:"tanggal"`
	Created         string `json:"created"`
}

type Total struct {
	JumlahPositif   int `json:"jumlah_positif"`
	JumlahMeninggal int `json:"jumlah_meninggal"`
	JumlahSembuh    int `json:"jumlah_sembuh"`
	JumlahDirawat   int `json:"jumlah_dirawat"`
}

var timezone *time.Location

func ParseUpdate(r io.Reader, source data.Source) (*UpdateResponse, error) {
	ur := UpdateResponse{Source: source}
	err := ur.Parse(r)
	if err != nil {
		return nil, err
	}
	return &ur, nil
}

func (ur *UpdateResponse) Parse(r io.Reader) error {
	return json.NewDecoder(r).Decode(ur)
}

func (ur *UpdateResponse) Normalize() (*data.DailySummary, error) {
	date, err := parseDate(ur.Update.Penambahan.Tanggal)
	if err != nil {
		return nil, err
	}
	updated, err := parseDateTime(ur.Update.Penambahan.Created)
	if err != nil {
		return nil, err
	}
	return &data.DailySummary{
		Metadata: data.Metadata{
			Country:   country.ID,
			Date:      date,
			UpdatedAt: updated,
			Source:    ur.Source,
		},
		Confirmed:           ur.Update.Penambahan.JumlahPositif,
		Recovered:           ur.Update.Penambahan.JumlahSembuh,
		Deceased:            ur.Update.Penambahan.JumlahMeninggal,
		ConfirmedCumulative: ur.Update.Total.JumlahPositif,
		RecoveredCumulative: ur.Update.Total.JumlahSembuh,
		DeceasedCumulative:  ur.Update.Total.JumlahMeninggal,
	}, nil
}

func parseDate(text string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", text, timezone)
}

func parseDateTime(text string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", text, timezone)
}

func init() {
	tz, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	timezone = tz
}
