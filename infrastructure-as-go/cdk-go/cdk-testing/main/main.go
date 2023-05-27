package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"cdktesting"
)

func main() {

	app := awscdk.NewApp(nil)

	cdktesting.NewCdkTestingStack(app, "CdkTestingStack", &cdktesting.CdkTestingStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	
	return nil

}
