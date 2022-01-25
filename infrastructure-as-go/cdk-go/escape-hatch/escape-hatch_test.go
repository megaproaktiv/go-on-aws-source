package escapehatch_test

import (
	"testing"

	"escapehatch"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestEscapeHatchStringStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := escapehatch.NewEscapeHatchStringStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"AnalyticsConfigurations": map[string]interface{}{
			"Id": "AnalyticsConfigurationId",
		},
	})
}
