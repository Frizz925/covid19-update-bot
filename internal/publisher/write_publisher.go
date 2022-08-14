package publisher

import (
	"fmt"
	"io"
	"time"
)

type WritePublisher struct {
	io.Writer
}

func NewWritePublisher(w io.Writer) *WritePublisher {
	return &WritePublisher{w}
}

func (wp *WritePublisher) Publish(message string) error {
	_, err := wp.Write([]byte(message + "\n"))
	return err
}

func (wp *WritePublisher) PublishEmbed(embed *Embed) error {
	var author, timestamp string
	if embed.Author.Name != "" {
		author = fmt.Sprintf("%s (%s)", embed.Author.Name, embed.Author.URL)
	}
	if !embed.Timestamp.IsZero() {
		timestamp = embed.Timestamp.Format(time.RFC3339)
	}

	messages := [][]string{
		{"Author", author},
		{"URL", embed.URL},
		{"Title", embed.Title},
		{"Description", embed.Description},
		{"ImageURL", embed.ImageURL},
		{"Footer", embed.Footer},
		{"Timestamp", timestamp},
	}

	messageCount := len(messages)
	fieldCount := len(embed.Fields)
	if fieldCount > 0 {
		fieldMessages := make([][]string, fieldCount)
		for idx, field := range embed.Fields {
			fieldMessages[idx] = []string{
				fmt.Sprintf("Field[%d]", idx),
				fmt.Sprintf("%s: %s", field.Name, field.Value),
			}
		}

		tmp := make([][]string, messageCount+fieldCount)
		copy(tmp, messages[:4])
		copy(tmp[4:], fieldMessages)
		copy(tmp[fieldCount+4:], messages[4:])
		messages = tmp
	}

	for _, tuple := range messages {
		key, message := tuple[0], tuple[1]
		if message == "" {
			continue
		}
		formatted := fmt.Sprintf("%s: %s", key, message)
		if err := wp.Publish(formatted); err != nil {
			return err
		}
	}
	return wp.Publish("")
}
