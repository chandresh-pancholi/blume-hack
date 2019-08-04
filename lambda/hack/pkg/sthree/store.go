package sthree

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type Store struct {
	S3 *s3.S3
}

func (s *Store) Get(bucket, key string) (io.ReadCloser, error) {
	output, err := s.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return output.Body, nil
}
