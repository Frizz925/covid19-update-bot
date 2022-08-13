package aws

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/frizz925/covid19-update-bot/internal/storage"
)

type S3Storage struct {
	*s3.S3
	region string
	bucket string
}

type s3Object struct {
	region string
	bucket string
	key    string
}

type s3Reader struct {
	s3Object
	io.ReadCloser
}

func NewS3Storage(sess *session.Session, region, bucket string) *S3Storage {
	return &S3Storage{s3.New(sess), region, bucket}
}

func (s *S3Storage) Read(ctx context.Context, name string) (storage.ObjectReader, error) {
	obj, err := s.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return nil, err
	}
	return &s3Reader{
		s3Object: s3Object{
			region: s.region,
			bucket: s.bucket,
			key:    name,
		},
		ReadCloser: obj.Body,
	}, nil
}

func (s *S3Storage) Write(ctx context.Context, name string, r io.Reader) (storage.ObjectFile, error) {
	var rs io.ReadSeeker
	if v, ok := r.(io.ReadSeeker); ok {
		rs = v
	} else {
		b, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		rs = bytes.NewReader(b)
	}
	req := s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
		Body:   rs,
	}
	if _, err := s.PutObjectWithContext(ctx, &req); err != nil {
		return nil, err
	}
	return &s3Object{
		region: s.region,
		bucket: s.bucket,
		key:    name,
	}, nil
}

func (o *s3Object) Name() string {
	return path.Join(o.bucket, o.key)
}

func (o *s3Object) URL() string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", o.bucket, o.region, o.key)
}
