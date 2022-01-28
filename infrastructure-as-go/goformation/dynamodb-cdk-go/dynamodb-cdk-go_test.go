package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/aws-sdk-go-v2/aws"
)

// example tests. To run these tests, uncomment this file along with the
// example resource in dynamodb-cdk-go_test.go
func TestDynamodbCdkGoStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewDynamodbCdkGoStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(aws.String("AWS::DynamoDB::Table"), map[string]interface{}{
		"TableName": "Username",
	})
}
