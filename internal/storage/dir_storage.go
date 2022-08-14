package storage

import (
	"context"
	"io"
	"path"
	"path/filepath"
)

type DirStorage struct {
	fsStorage
	dir string
}

func NewDirStorage(dirs ...string) *DirStorage {
	return &DirStorage{
		dir: path.Join(dirs...),
	}
}

func (s *DirStorage) Read(_ context.Context, name string) (ObjectReader, error) {
	p, err := s.getPath(name)
	if err != nil {
		return nil, err
	}
	return s.read(p)
}

func (s *DirStorage) Write(_ context.Context, name string, r io.Reader) (ObjectFile, error) {
	p, err := s.getPath(name)
	if err != nil {
		return nil, err
	}
	return s.write(p, r)
}

func (s *DirStorage) getPath(name string) (string, error) {
	return filepath.Abs(path.Join(s.dir, name))
}
