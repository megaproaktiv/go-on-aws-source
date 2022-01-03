package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"cfnresource"
)


func main() {
	app := awscdk.NewApp(nil)

	cfnresource.NewCfnresourceStack(app, "CfnresourceStack", &cfnresource.CfnresourceStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
