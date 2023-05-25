package gograviton

import (
	"os"
	"log"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/constructs-go/constructs/v10"
)

type LambdaGoArmStackProps struct {
	awscdk.StackProps
}

func NewLambdaGoArmStack(scope constructs.Construct, id string, props *LambdaGoArmStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	dockerfile := filepath.Join(path, "../app")
	//begin lambda_arm
	awslambda.NewDockerImageFunction(stack,
		aws.String("RegisterHandlerArm"),
		&awslambda.DockerImageFunctionProps{
			Architecture:                 awslambda.Architecture_ARM_64(),
			FunctionName:                 aws.String("hellodockerarm"),
			MemorySize:                   aws.Float64(1024),
			Timeout:                      awscdk.Duration_Seconds(aws.Float64(300)),
			Code:                         awslambda.DockerImageCode_FromImageAsset(&dockerfile, &awslambda.AssetImageCodeProps{}),
		})
	//end lambda_arm
	dockerfile = filepath.Join(path, "../appx86")
	awslambda.NewDockerImageFunction(stack,
		aws.String("RegisterHandlerAmd"),
		&awslambda.DockerImageFunctionProps{
			Architecture:                 awslambda.Architecture_X86_64(),
			FunctionName:                 aws.String("hellodockerx86"),
			MemorySize:                   aws.Float64(1024),
			Timeout:                      awscdk.Duration_Seconds(aws.Float64(300)),
			Code:                         awslambda.DockerImageCode_FromImageAsset(&dockerfile, &awslambda.AssetImageCodeProps{}),
		})


	return stack
}
