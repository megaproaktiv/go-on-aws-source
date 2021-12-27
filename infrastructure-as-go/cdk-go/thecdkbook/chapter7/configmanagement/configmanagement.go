package simpleapiwithtestsstack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type SimpleApiWithTestsStackProps struct {
	awscdk.StackProps
}

func NewSimpleApiWithTestsStack(scope constructs.Construct, id string, props *SimpleApiWithTestsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	return stack
}
