package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type VpcStackProps struct {
	awscdk.StackProps
}

func NewVpcStack(scope constructs.Construct, id string, props *VpcStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	myVpc := awsec2.NewVpc(stack, aws.String("basevpc"),
		&awsec2.VpcProps{
			Cidr: aws.String("10.0.0.0/16"),
			MaxAzs: aws.Float64(1),
		},
	)

	awsssm.NewStringParameter(stack, aws.String("basevpc-parm"),
		&awsssm.StringParameterProps{
			Description:    aws.String("Created VPC"),
			ParameterName:  aws.String("/network/basevpc"),
			StringValue:    myVpc.VpcId(),
		},
	)

	return stack
}
