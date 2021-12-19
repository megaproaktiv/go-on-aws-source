package customresource_test

import (
	"customresource"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)


func TestCustomresourceStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := customresource.NewCustomresourceStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(jsii.String("AWS::CloudFormation::CustomResource"), map[string]interface{}{
		"ServiceToken":  map[string]interface{}{
			"Fn::GetAtt": []string{
				"myHandler0D56A5FA",
				"Arn",
			},
		},
	})
}
