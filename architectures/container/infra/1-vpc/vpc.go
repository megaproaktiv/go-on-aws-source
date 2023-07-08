package vpc

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type VpcProps struct {
	cdk.StackProps
}

func VpcStack(scope constructs.Construct, id string, props *VpcProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)
	vpc := awsec2.NewVpc(stack, aws.String("go-on-aws-vpc"), &awsec2.VpcProps{
		NatGateways: aws.Float64(1),
	})
	//begin parameter
	vpcParm := "/go-on-aws/vpc"
	awsssm.NewStringParameter(stack, &vpcParm,
		&awsssm.StringParameterProps{
			AllowedPattern: new(string),
			Description:    aws.String("VPC id go on aws architecture"),
			ParameterName:  &vpcParm,
			StringValue:    vpc.VpcId(),
		})
	//end parameter
	return stack
}
