package publisher

import "time"

type Publisher interface {
	Publish(message string) error
	PublishEmbed(embed *Embed) error
}

type Embed struct {
	URL         string
	Author      Author
	Title       string
	Description string
	Fields      []Field
	ImageURL    string
	Footer      string
	Timestamp   time.Time
}

type Author struct {
	Name string
	URL  string
}

type Field struct {
	Name  string
	Value string
}
