package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)


func main() {
	app := awscdk.NewApp(nil)

	parentStack := awscdk.NewStack(app, aws.String("MyParentStack"), nil)

	awscdk.NewNestedStack(parentStack, aws.String("MyNestedStack"),nil)	
	
	app.Synth(nil)
}
