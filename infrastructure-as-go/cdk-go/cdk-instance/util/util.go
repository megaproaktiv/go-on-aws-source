package util

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)



func Env() *awscdk.Environment {
	return &awscdk.Environment{
	 Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}