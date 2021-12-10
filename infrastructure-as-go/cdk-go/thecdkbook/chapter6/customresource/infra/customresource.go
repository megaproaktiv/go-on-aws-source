package customresource

import (
	"log"
	"os"
	"path/filepath"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	logs "github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"	

)

type CustomresourceStackProps struct {
	awscdk.StackProps
}

func NewCustomresourceStack(scope constructs.Construct, id string, props *CustomresourceStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	this := stack
	// The code that defines your stack goes here
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../app/dist/main.zip")

	myFunction := awslambda.NewFunction(stack, aws.String("myHandler"), 
	&awslambda.FunctionProps{
		Description:                  aws.String("customresourcelambda"),
		FunctionName:                 aws.String("customresourcelambda"),
		LogRetention:                 logs.RetentionDays_THREE_MONTHS,
		MemorySize:                   aws.Float64(1024),
		Timeout:                      awscdk.Duration_Seconds(aws.Float64(10)),
		Code: awslambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
		Handler: aws.String("main"),
		Runtime: awslambda.Runtime_GO_1_X(),
	})

	awscdk.NewCustomResource(this, aws.String("MyResource"), &awscdk.CustomResourceProps{
		ServiceToken:         myFunction.FunctionArn(),
		
	})
	
	

	return stack
}
