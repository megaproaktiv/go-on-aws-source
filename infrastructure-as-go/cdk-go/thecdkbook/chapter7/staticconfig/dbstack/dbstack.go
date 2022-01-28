package dbstack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	
)

type DbstackStackProps struct {
	awscdk.StackProps
	VpcID *string
	InstanceType *string
}

func NewDbstackStack(scope constructs.Construct, id string, props *DbstackStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	return stack
}
