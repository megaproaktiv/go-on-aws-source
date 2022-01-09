package dockerbundler_test

import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	
	"dockerbundler"
)


func TestDockerbundlerStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := dockerbundler.NewDockerbundlerStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(aws.String("AWS::Lambda::Function"), map[string]interface{}{
		"Runtime": "go1.x",
	})
}
