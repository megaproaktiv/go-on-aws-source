package s3share

//begin import
import (
	"context"
	"log"
	"time"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)
//end import
// Share existing S3 object with presigned URL
// begin function
func Share(client *s3.Client, key *string, bucket *string) (string, error) {
	//end function

	//begin return
	var url string
	//end return
	// generate the presigned URL
	// s3 presign client
	//begin presign
	// Set the expiration time for the presigned URL
	lifetimeSecs := int64(3600)
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
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			*bucket, *key, err)
		return "", err
	}
	//end presign
	//begin return
	url = string(req.URL)
	//end return

	//begin function
	return url, nil
}

//end function
