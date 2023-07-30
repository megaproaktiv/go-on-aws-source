package hello

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	logs "github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	ssm "github.com/aws/aws-cdk-go/awscdk/v2/awsssm"

	"github.com/aws/constructs-go/constructs/v10"
)

type HelloWorldStackProps struct {
	StackProps awscdk.StackProps
}

func NewHelloWorldStack(scope constructs.Construct, id string, props *HelloWorldStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../dist/bootstrap.zip")

	//begin lambda
	fn := lambda.NewFunction(stack, aws.String("simplelambda"), 
	&lambda.FunctionProps{
		Description:                  aws.String("Simple Lambda says hello"),
		FunctionName:                 aws.String("sayhello"),
		LogRetention:                 logs.RetentionDays_THREE_MONTHS,
		MemorySize:                   aws.Float64(1024),
		Timeout:                      awscdk.Duration_Seconds(aws.Float64(3)),
		Code: lambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
		Handler:      aws.String("bootstrap"),
		Runtime:      lambda.Runtime_PROVIDED_AL2(),
		Architecture: lambda.Architecture_ARM_64(),
		
	})
	//end lambda

	ssm.NewStringParameter(stack, aws.String("Functionname"), &ssm.StringParameterProps{
		Description:    aws.String("helloworldlamba"),
		ParameterName:  aws.String("simplefunction"),
		StringValue:    fn.FunctionName(),
	})

	return stack
}
