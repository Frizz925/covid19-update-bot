package storage

import (
	"context"
	"io"
	"os"
	"path"
)

type TempStorage struct {
	fsStorage
	dir string
}

func NewTempStorage(dirs ...string) *TempStorage {
	return &TempStorage{
		dir: path.Join(dirs...),
	}
}

func (s *TempStorage) Read(_ context.Context, name string) (ObjectReader, error) {
	return s.read(s.getPath(name))
}

func (s *TempStorage) Write(_ context.Context, name string, r io.Reader) (ObjectFile, error) {
	return s.write(s.getPath(name), r)
}

func (s *TempStorage) getPath(name string) string {
	return path.Join(os.TempDir(), s.dir, name)
}
