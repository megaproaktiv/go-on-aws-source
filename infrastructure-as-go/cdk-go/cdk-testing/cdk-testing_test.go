package cdktesting_test

import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	//begin import
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	//end  import
	"github.com/aws/jsii-runtime-go"

	"cdktesting"
)

//begin test
func TestCdkTestingStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := cdktesting.NewCdkTestingStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack,nil)
	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"VisibilityTimeout": 300,
	})
	template.ResourceCountIs(jsii.String("AWS::SQS::Queue"), aws.Float64(1))

}
//end test
