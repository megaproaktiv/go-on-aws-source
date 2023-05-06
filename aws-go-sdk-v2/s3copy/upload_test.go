package s3copy_test

import (
	"context"
	"testing"

	"s3copy"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/megaproaktiv/awsmock"
	"gotest.tools/assert"
)

func TestUpload(t *testing.T) {
	mock := func(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
		out := &s3.PutObjectOutput{
			ETag:                   aws.String("\"d41d8cd98f00b204e9800998ecf8427e\""),
			ServerSideEncryption:    "AES256",
		}
		assert.Equal(t, *params.Bucket, "testb")
		assert.Equal(t, *params.Key, "testp/test.txt")
		return out, nil
	}

	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(mock)
	client := s3.NewFromConfig(mockCfg.AwsConfig())

	file := "testdata/test.txt"
	bucket := "testb"
	prefix := "testp/"
	err := s3copy.Upload(client,  &file, &bucket, &prefix)
	assert.NilError(t, err, "should give no error")

}
