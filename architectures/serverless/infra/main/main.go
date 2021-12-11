package main

import (
	"dsl"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	app := awscdk.NewApp(nil)

	dsl.NewDslStack(app, "dsl", &dsl.DslStackProps{
		awscdk.StackProps{
			Env: env(),
			Description: aws.String("Serverless Standard GOA Demo"),
		},
	})

	app.Synth(nil)
}


func env() *awscdk.Environment {
	return nil
}
