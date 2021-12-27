package dbstack_test

import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"dbstack"
)

func TestDbstackStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := dbstack.NewDbstackStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.ResourceCountIs( aws.String("AWS::CDK::Metadata"), aws.Float64(0))
}
