package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load AWS Credentials from .env
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_S3_BUCKET")
	uploadDir := os.Getenv("UPLOAD_DIR") // Load upload directory from .env

	if uploadDir == "" {
		log.Fatal("UPLOAD_DIR is not set in .env")
	}

	// Load AWS SDK Config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	// Walk through all files in UPLOAD_DIR
	err = filepath.WalkDir(uploadDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() { // Upload only files, skip directories
			relativePath, _ := filepath.Rel(uploadDir, path)
			err := uploadToS3(s3Client, bucketName, path, relativePath)
			if err != nil {
				log.Printf("Failed to upload %s: %v", path, err)
			} else {
				log.Printf("Uploaded: %s", relativePath)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}
}

func uploadToS3(client *s3.Client, bucket, filePath, s3Key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &s3Key,
		Body:   file,
	})
	return err
}
