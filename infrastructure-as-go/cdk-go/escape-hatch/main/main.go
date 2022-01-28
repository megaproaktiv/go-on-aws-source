package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"escapehatch"
)

func main() {
	app := awscdk.NewApp(nil)
	escapehatch.NewEscapeHatchStringStack(app, "string", &escapehatch.EscapeHatchStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	},)

	escapehatch.NewEscapeHatchFileStack(app, "file", &escapehatch.EscapeHatchStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	},)

	escapehatch.NewEscapeHatchStringStack(app, "struct", &escapehatch.EscapeHatchStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	},)

	app.Synth(nil)
}


func env() *awscdk.Environment {
	return nil
}
