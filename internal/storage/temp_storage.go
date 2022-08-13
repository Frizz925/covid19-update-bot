package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
)

type TempStorage struct{}

type tempObject struct {
	name string
	url  string
}

type tempReader struct {
	tempObject
	io.ReadCloser
}

func NewTempStorage() *TempStorage {
	return &TempStorage{}
}

func (t *TempStorage) Read(ctx context.Context, name string) (ObjectReader, error) {
	f, err := os.Open(t.getPath(name))
	if err != nil {
		return nil, err
	}
	return &tempReader{
		tempObject: tempObject{
			name: f.Name(),
			url:  t.getURL(f.Name()),
		},
		ReadCloser: f,
	}, nil
}

func (t *TempStorage) Write(ctx context.Context, name string, r io.Reader) (ObjectFile, error) {
	dirname, basename := t.getPath(path.Dir(name)), path.Base(name)
	if _, err := os.Stat(dirname); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		if err := os.MkdirAll(dirname, 0755); err != nil {
			return nil, err
		}
	}
	f, err := os.CreateTemp(dirname, "*-"+basename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(f, r); err != nil {
		return nil, err
	}
	return &tempObject{
		name: f.Name(),
		url:  t.getURL(f.Name()),
	}, nil
}

func (TempStorage) getPath(name string) string {
	return path.Join(os.TempDir(), name)
}

func (TempStorage) getURL(name string) string {
	return fmt.Sprintf("file://%s", name)
}

func (o *tempObject) Name() string {
	return o.name
}

func (o *tempObject) URL() string {
	return o.url
}
