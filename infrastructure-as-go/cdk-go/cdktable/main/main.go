package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/jsii-runtime-go"
	"cdktable"
)

func main() {
	// defer jsii.Close()

	app := awscdk.NewApp(nil)

	//begin createstack
	cdktable.NewCdktableStack(app, "table", &cdktable.CdktableStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	},
	)
	//end createstack

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
