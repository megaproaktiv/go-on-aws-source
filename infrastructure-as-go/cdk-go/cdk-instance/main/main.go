package main

import (
	"instance"
	"instance/util"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := awscdk.NewApp(nil)

	instance.NewInstanceStack(app, "instance", &instance.InstanceStackProps{
		StackProps: awscdk.StackProps{
			Env: util.Env(),
		},
	})

	app.Synth(nil)
}
