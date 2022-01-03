package cfnresource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-sdk-go-v2/aws"

)

type CfnresourceStackProps struct {
	awscdk.StackProps
}

func NewCfnresourceStack(scope constructs.Construct, id string, props *CfnresourceStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	this := stack

	awscdk.NewCfnResource(this,aws.String("MyResource"), &awscdk.CfnResourceProps{
		Properties: &map[string]interface{}{
			"Name": "MyResource",
			"Enabled" : aws.Bool(true),
		},
		Type:       aws.String("Vendor::My::Resource"),
	} )

	return stack
}
