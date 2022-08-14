package storage

import (
	"fmt"
	"io"
	"os"
	"path"
)

type fsStorage struct{}

type fsObject struct {
	name string
	url  string
}

type fsReader struct {
	fsObject
	io.ReadCloser
}

func (s *fsStorage) read(filename string) (ObjectReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &fsReader{
		fsObject: fsObject{
			name: f.Name(),
			url:  s.getURL(f.Name()),
		},
		ReadCloser: f,
	}, nil
}

func (s *fsStorage) write(filename string, r io.Reader) (ObjectFile, error) {
	dirname, basename := path.Dir(filename), path.Base(filename)
	if _, err := os.Stat(dirname); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		if err := os.MkdirAll(dirname, 0755); err != nil {
			return nil, err
		}
	}
	f, err := os.Create(path.Join(dirname, basename))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(f, r); err != nil {
		return nil, err
	}
	return &fsObject{
		name: f.Name(),
		url:  s.getURL(f.Name()),
	}, nil
}

func (fsStorage) getURL(name string) string {
	return fmt.Sprintf("file://%s", name)
}

func (o *fsObject) Name() string {
	return o.name
}

func (o *fsObject) URL() string {
	return o.url
}
