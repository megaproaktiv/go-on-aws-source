package s3share

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Share existing S3 object with presigned URL
func Share(client *s3.Client, key *string, bucket *string) (string, error) {

	var url string
	// Set the expiration time for the presigned URL
	lifetimeSecs := int64(3600)

	// generate the presigned URL
	// s3 presign client
	//begin presign
	s3PresignClient := s3.NewPresignClient(client)
	req, err := s3PresignClient.PresignGetObject(
		context.TODO(),
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
		})
	//end presign
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			*bucket, *key, err)
		return "", err
	}
	url = string(req.URL)

	return url, nil
}
