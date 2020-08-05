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
	BucketName := os.Getenv("BUCKET")
	orgObjectKey := os.Getenv("ORG_OBJECT_KEY")
	newObjectKey := os.Getenv("NEW_OBJECT_KEY")

	// Create session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "s3",
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 client
	svc := s3.New(sess)

	// Move Data
	_, err = svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(BucketName),
		CopySource: aws.String(BucketName + "/" + orgObjectKey),
		Key:        aws.String(newObjectKey)})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("file move done")
}
