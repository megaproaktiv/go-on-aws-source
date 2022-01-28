package simpleapiwithtestsstack_test

import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	
	"simpleapiwithtestsstack"
)

// example tests. To run these tests, uncomment this file along with the
// example resource in configmanagement_test.go
func TestConfigmanagementStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack :=simpleapiwithtestsstack.NewSimpleApiWithTestsStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.ResourceCountIs( aws.String("AWS::CDK::Metadata"), aws.Float64(0))
}
