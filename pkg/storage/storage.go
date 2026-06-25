package storage

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"go-starter-backend/internal/config"
)

type Storage struct {
	client    *s3.Client
	presigner *s3.PresignClient
	bucket    string
	publicURL string
}

func New(cfg *config.Config) (*Storage, error) {
	awsCfg, err := awsConfig.LoadDefaultConfig(
		context.Background(),
		awsConfig.WithRegion(cfg.S3Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.S3AccessKey,
				cfg.S3SecretKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.S3Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.S3Endpoint)
		}

		o.UsePathStyle = cfg.S3ForcePathStyle
	})

	return &Storage{
		client:    client,
		presigner: s3.NewPresignClient(client),
		bucket:    cfg.S3Bucket,
		publicURL: strings.TrimRight(cfg.S3PublicURL, "/"),
	}, nil
}

func (s *Storage) Upload(ctx context.Context, key string, body io.Reader, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})

	return err
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	return err
}

func (s *Storage) PublicURL(key string) string {
	if s.publicURL == "" {
		return ""
	}

	key = strings.TrimLeft(key, "/")

	return fmt.Sprintf("%s/%s", s.publicURL, key)
}

func (s *Storage) PresignedURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	result, err := s.presigner.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(key),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = expires
		},
	)

	if err != nil {
		return "", err
	}

	return result.URL, nil
}
