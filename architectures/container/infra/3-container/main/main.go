package main

import (
	"os"
	"showtable"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	app := awscdk.NewApp(nil)

	fsp := new(showtable.FargateStackStackProps)
	fsp.Env = env()
	showtable.FargateStack(app, "cluster", fsp)

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region:  aws.String("eu-central-1"),
		Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	}

}
