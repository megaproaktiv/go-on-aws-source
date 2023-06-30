package s3share

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Load file to s3
func Upload(client *s3.Client, filename *string,bucket *string) error {
	
	//begin upload
	content, err := os.ReadFile(*filename)

	//end upload
	//handle error
	if err != nil {
		return err
	}

	// Upload  file to s3 
	//begin upload
	key := filename
	_, err = client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
			Body:   bytes.NewReader(content),
		}, 
	)
	//end upload
	//handle error
	if err != nil {
		return err
	}


	return nil
}