package publisher

type Publisher interface {
	Publish(message string) error
	PublishEmbed(embed *Embed) error
}

type Embed struct {
	Title    string
	Content  string
	ImageURL string
	URL      string
}
