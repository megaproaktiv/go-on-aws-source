package main

import (
	"bufio"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	fromPtr := flag.String("from", "", "From source")
	flag.Parse()

	fromPtr := flag.String("from", "", "From source file")
	bucketPtr := flag.String("bucket", "", "sharing bucket")

	flag.Parse()

	s3share.Upload(fromPtr, bucketPtr, s3share.Client)

		u, err := url.Parse(*toPtr)

		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(*fromPtr)

		if err != nil {
			log.Fatal(err)

			return
		}

		uploader := s3manager.NewUploader(session.New())

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(u.Host),
			Key:    aws.String(u.Path),
			Body:   aws.ReadSeekCloser(bufio.NewReader(f)),
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					log.Fatal(aerr)
				}
			} else {
				log.Fatal(err)
			}

			return
		}

}
