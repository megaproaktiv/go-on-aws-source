package s3share_test

import (
	"context"
	"testing"

	"github.com/megaproaktiv/awsmock"
	"gotest.tools/assert"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"s3share"
)

func TestUpload(t *testing.T) {
	callCount := 0
	putObjectMock := func(ctx context.Context, params *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
		callCount++
		return &s3.PutObjectOutput{}, nil
	}
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(putObjectMock)
	client := s3.NewFromConfig(mockCfg.AwsConfig())
	err := s3share.Upload(client, aws.String("testdata/text.txt"), aws.String("bucket"))
	assert.NilError(t, err)
}