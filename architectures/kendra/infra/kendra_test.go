package kendra_test

import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
)


func TestKendraStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewKendraStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(aws.String("AWS::Kendra::Index"), map[string]interface{}{
		"Edition": "DEVELOPER_EDITION",
	})
}
