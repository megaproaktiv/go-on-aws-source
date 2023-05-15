package s3share

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Load file to s3
func Upload(client *s3.Client, filename *string,bucket *string) error {
	
	content, err := os.ReadFile(*filename)
	//handle error
	if err != nil {
		return err
	}

	key := filename
	// Upload  file to s3 
	_, err = client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
			Body:   bytes.NewReader(content),
		}, 
	)
	//handle error
	if err != nil {
		return err
	}


	return nil
}