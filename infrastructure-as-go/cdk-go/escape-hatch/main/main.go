package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"escapehatch"
)

func main() {
	app := awscdk.NewApp(nil)
	escapehatch.NewEscapeHatchStack(app, "EscapeHatchStack", &escapehatch.EscapeHatchStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	},)

	app.Synth(nil)
}


func env() *awscdk.Environment {
	return nil
}
