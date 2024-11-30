package r2

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

// CreateS3Client initializes the S3 client
func CreateS3Client() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3endpoint := os.Getenv("AWS_S3_ENDPOINT")
	if s3endpoint == "" {
		return errors.New("missing environment variable: AWS_S3_ENDPOINT")
	}

	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(s3endpoint)
	})

	return nil
}

// UploadFile uploads a file to the specified S3 bucket
func UploadFile(objectKey string, r io.Reader) error {
	if client == nil {
		return errors.New("S3 client is not initialized")
	}

	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	if bucketName == "" {
		return errors.New("missing environment variable: AWS_S3_BUCKET_NAME")
	}


	// ここでobjectKeyの拡張子を取得して、それに応じたContent-Typeを設定する
	objectKeyParts := strings.Split(objectKey, ".")
	if len(objectKeyParts) < 2 {
		return fmt.Errorf("invalid object key: %s", objectKey)
	}
	ext := "." + objectKeyParts[len(objectKeyParts)-1]
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        r,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return nil
}

// URL生成

func GenerateURL(objectKey string) (string, error) {

		// R2オブジェクトのURL生成
	buketUrl := os.Getenv("AWS_S3_BUCKET_URL")

	if buketUrl == "" {
		return "", errors.New("missing environment variable: AWS_S3_BUCKET_URL")
	}

		imageURL := fmt.Sprintf("%s/%s",buketUrl,objectKey)

	return imageURL, nil
	
}
