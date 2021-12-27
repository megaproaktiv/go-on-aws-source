package main

import (
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"simpleapiwithtestsstack"
)

func main() {
	devEnv := &awscdk.Environment{
		Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region: aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
	
	app := awscdk.NewApp(nil)

	simpleapiwithtestsstack.NewSimpleApiWithTestsStack(app, "SimpleApiWithTests", 
		&simpleapiwithtestsstack.SimpleApiWithTestsStackProps{
			StackProps: awscdk.StackProps{
				Env: devEnv,
			},
		},
	)

	app.Synth(nil)
}

