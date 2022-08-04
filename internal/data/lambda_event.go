package data

type LambdaEvent struct {
	CountryID string               `json:"country_id"`
	Channels  []LambdaEventChannel `json:"channels"`
}

type LambdaEventChannel struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
}
