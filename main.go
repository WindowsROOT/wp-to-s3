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
	// โหลด AWS Credentials จาก .env
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_S3_BUCKET")

	// กำหนดโฟลเดอร์ที่ต้องการอัปโหลด
	uploadDir := "uploads"

	// โหลด AWS SDK
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	// สร้าง S3 Client
	s3Client := s3.NewFromConfig(cfg)

	// วนลูปหาไฟล์ทั้งหมดในโฟลเดอร์ uploads
	err = filepath.WalkDir(uploadDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() { // อัปโหลดเฉพาะไฟล์ (ไม่รวมโฟลเดอร์)
			relativePath, _ := filepath.Rel(uploadDir, path) // เอา path โดยไม่รวม "uploads/"
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

// ฟังก์ชันอัปโหลดไฟล์ไป S3
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
