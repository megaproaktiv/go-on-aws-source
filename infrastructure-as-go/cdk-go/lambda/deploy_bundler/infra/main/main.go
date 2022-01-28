package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"dockerbundler"
)

func main() {
	app := awscdk.NewApp(nil)

	dockerbundler.NewDockerbundlerStack(app, "DockerbundlerStack", &dockerbundler.DockerbundlerStackProps{
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

	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
