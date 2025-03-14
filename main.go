package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	godotenv.Load()

	// Load AWS Credentials from .env
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_S3_BUCKET")
	uploadDir := os.Getenv("UPLOAD_DIR")

	if uploadDir == "" {
		log.Fatal("UPLOAD_DIR is not set in .env")
	}

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create S3 client
	s3Client := s3.New(sess)

	// Walk through all files in UPLOAD_DIR
	err = filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() { // Upload only files, skip directories
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

func uploadToS3(client *s3.S3, bucket, filePath, s3Key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Key),
		Body:   file,
	})
	return err
}
