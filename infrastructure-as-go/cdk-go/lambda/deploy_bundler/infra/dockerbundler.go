package dockerbundler

import (
	"github.com/aws/aws-sdk-go-v2/aws"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	ssm "github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	asset "github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

type DockerbundlerStackProps struct {
	cdk.StackProps
}

func NewDockerbundlerStack(scope constructs.Construct, id string, props *DockerbundlerStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)
	this := stack

	
	handler := lambda.NewFunction(this, aws.String("GoLangDockerECRImageBundle"), 
		&lambda.FunctionProps{
			Handler:                      aws.String("main"),
			Runtime:                      lambda.Runtime_GO_1_X(),
			Code:                         lambda.AssetCode_FromAsset(aws.String("../app"), &asset.AssetOptions{
				Bundling:       &cdk.BundlingOptions{
					Image:            lambda.Runtime_GO_1_X().BundlingImage(),
					User:             aws.String("root"),
					Command:          &[]*string{
						aws.String("bash"),
						aws.String("-c"),
						aws.String("go version && go build -o /asset-output/main"),
					},
				},
			}),
		},
	)

	ssm.NewStringParameter(stack, aws.String("bundlefunction"), &ssm.StringParameterProps{
		Description:    aws.String("bundlefunction"),
		ParameterName:  aws.String("bundlefunction"),
		StringValue:    handler.FunctionName(),
	  })
	

	return stack
}
