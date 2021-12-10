package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)


func main() {
	app := awscdk.NewApp(nil)

	// Go does not allow unsed variables
	awscdk.NewStack(app, aws.String("MyFrontend"), nil)
	awscdk.NewStack(app, aws.String("MyBackend"), nil)
	
	app.Synth(nil)
}
