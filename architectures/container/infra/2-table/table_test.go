package table_test


import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	
	"table"
)

func TestTableStack(t *testing.T){
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := table.TableStack(app, "myStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)


	template.HasResource(aws.String("AWS::DynamoDB::Table"), 
		map[string]interface{}{
			"UpdateReplacePolicy": "Delete",
		},
	)
	// // To Test Fail
	// template.HasResourceProperties(aws.String("AWS::DynamoDB::Table"), 
	// 	map[string]interface{}{
	// 		"UpdateReplacePolicy": "Delete",
	// 	},
	// )
}