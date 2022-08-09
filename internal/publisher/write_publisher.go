package publisher

import "io"

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
