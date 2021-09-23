package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"instance"
	"instance/util"
)

func main() {
	app := awscdk.NewApp(nil)

	instance.NewInstanceStack(app, "instance", &instance.InstanceStackProps{
		awscdk.StackProps{
			Env: util.Env(),
		},
	})

	app.Synth(nil)
}


