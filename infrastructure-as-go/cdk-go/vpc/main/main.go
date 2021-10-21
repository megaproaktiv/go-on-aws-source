package main

import (
	"vpc"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	app := awscdk.NewApp(nil)

	vpc.NewVpcStack(app, "basevpc", &vpc.VpcStackProps{
		awscdk.StackProps{
			Env: env(),
			Description: aws.String("Base vpc"),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil

}
