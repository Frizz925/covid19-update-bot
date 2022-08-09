package data

type LambdaEvent struct {
	DataSources map[string]string    `json:"data_sources"`
	Channels    []LambdaEventChannel `json:"channels"`
}

type LambdaEventChannel struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
}
