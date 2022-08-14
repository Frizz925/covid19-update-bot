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
	if err := wp.Publish(embed.URL); err != nil {
		return err
	}
	if err := wp.Publish(embed.Title); err != nil {
		return err
	}
	if err := wp.Publish(embed.Description); err != nil {
		return err
	}
	for _, field := range embed.Fields {
		if err := wp.Publish(fmt.Sprintf("%s: %s", field.Name, field.Value)); err != nil {
			return err
		}
	}
	if err := wp.Publish(embed.ImageURL); err != nil {
		return err
	}
	if err := wp.Publish(embed.Footer); err != nil {
		return err
	}
	return wp.Publish(embed.Timestamp.Format(time.RFC3339))
}
