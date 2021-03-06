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

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(dlBucketName),
		Key:    aws.String(dlObjectKey),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctype := *obj.ContentType
	date := obj.LastModified
	log.Printf("content type: %s, last modified: %v", ctype, date)

	// read download file only 10 bytes
	rc := obj.Body
	defer rc.Close()
	buf := make([]byte, 10)
	_, err = rc.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
}
