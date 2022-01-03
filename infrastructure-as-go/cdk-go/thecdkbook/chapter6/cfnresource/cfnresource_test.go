package cfnresource_test

import (
	"testing"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/aws-sdk-go-v2/aws"
	"cfnresource"
)


func TestCfnresourceStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := cfnresource.NewCfnresourceStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(aws.String("Vendor::My::Resource"), map[string]interface{}{
		"Enabled": aws.Bool(true),
	})
}
