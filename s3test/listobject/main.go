package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	upBucketName := os.Getenv("UP_BUCKET")
	searchString := os.Getenv("STR")

	// Create session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "s3",
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 client
	svc := s3.New(sess)

	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(upBucketName)})
	if err != nil {
		log.Println(err)
	}

	for _, item := range resp.Contents {
		if strings.Index(*item.Key, searchString) != -1 {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
}
