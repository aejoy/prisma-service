package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Storage struct {
	domain  string
	bucket  string
	session *s3.Client
}

func NewStorage(domain, endpoint, key, secret, bucket, region string) (*Storage, error) {
	credentialsProvider := config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(key, secret, ""),
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		credentialsProvider, config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	session := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.BaseEndpoint = aws.String(endpoint)
	})

	return &Storage{domain, bucket, session}, nil
}

func (s *Storage) SaveObject(key string, body io.Reader) (string, error) {
	_, err := s.session.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String("image/avif"),
		ACL:         types.ObjectCannedACLPrivate,
	})

	return fmt.Sprintf("https://%s/%s", s.domain, key), err
}
