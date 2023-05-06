package s3copy

// Import go sdk v2 s3 and config

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var Client *s3.Client

// Init the s3 client
func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("Error loading AWS config: %s", err)
		os.Exit(1)
	}
	Client = s3.NewFromConfig(cfg)
}

// Take a file as parameter and upload it to the s3 bucket

func Upload(client *s3.Client,file *string, bucket *string, prefix *string) error {
	// Open the file
	//begin open
	body, err := os.ReadFile(*file)
	if err != nil {
		fmt.Printf("File: %v could not be opened\n",err)
		return err
	}
	//end open

	// Upload the file
	key := *prefix + path.Base(*file)
	//begin upload
	parms := &s3.PutObjectInput{
		Bucket: bucket,
		Key: aws.String(key),
		Body: bytes.NewReader(body),
	}
	_, err = client.PutObject(context.TODO(), parms)
	//end upload
	//begin error
	if err != nil {
		fmt.Printf("Error uploading to S3 Bucket: %v Key: %v\n",bucket,key)
		return err
	}
	//end erro
	
	return nil

}