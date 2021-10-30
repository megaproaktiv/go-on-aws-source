package main

import (
	"gograviton"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)


func main() {
	app := awscdk.NewApp(nil)
	gograviton.NewLambdaGoArmStack(app, "lambda-go-arm", &gograviton.LambdaGoArmStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}


func env() *awscdk.Environment {
	return nil

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
