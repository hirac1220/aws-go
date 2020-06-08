package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	upBucketName := os.Getenv("UP_BUCKET")
	upObjectKey := os.Getenv("UP_OBJECT_KEY")

	// Create session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "s3",
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Read file
	targetFilePath := "./data/upload.png"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Make Uploader
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(upBucketName),
		Key:    aws.String(upObjectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("file upload done")
}
