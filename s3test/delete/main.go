package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dlBucketName := os.Getenv("DL_BUCKET")
	dlObjectKey := os.Getenv("DL_OBJECT_KEY")

	// Create session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "s3",
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 client
	svc := s3.New(sess)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(dlBucketName), Key: aws.String(dlObjectKey)})
	if err != nil {
		log.Println(err)
	}
	log.Println("file deleted")
}
