package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	// 2.3.1. The Stack Construct
	app := awscdk.NewApp(nil)
	parentStack := awscdk.NewStack(app, aws.String("MyParentStack"), nil)
	awscdk.NewStack(parentStack, aws.String("MyChildStack"), nil)

	app.Synth(nil)
}
