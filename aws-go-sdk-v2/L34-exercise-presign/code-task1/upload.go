package s3share

//file: upload.go
import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Load file to s3
// begin function
func Upload(client *s3.Client, filename *string, bucket *string) error {
	//end function

	//begin upload
	content, err := os.ReadFile(*filename)

	//end upload
	//handle error
	//begin error
	if err != nil {
		return err
	}
	//end  error

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

	//begin function
	return nil
}

//end function
