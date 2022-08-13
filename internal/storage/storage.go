package storage

import (
	"context"
	"io"
)

type Type string

const (
	Temp Type = "tmp"
)

type Storage interface {
	Read(ctx context.Context, name string) (ObjectReader, error)
	Write(ctx context.Context, name string, r io.Reader) (ObjectFile, error)
}

type ObjectFile interface {
	Name() string
	URL() string
}

type ObjectReader interface {
	ObjectFile
	io.ReadCloser
}
