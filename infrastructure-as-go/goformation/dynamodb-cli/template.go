package gfcli

import (
	"github.com/awslabs/goformation/v5/cloudformation"
	"github.com/awslabs/goformation/v5/cloudformation/dynamodb"
	"github.com/awslabs/goformation/v5/cloudformation/tags"
)

// CreateTemplate - build the Cloudformation template
func CreateTemplate(name string) (*cloudformation.Template, error) {
	template := cloudformation.NewTemplate()

	// Create an the Table
	template.Resources["simpletable"] = &dynamodb.Table{
		AttributeDefinitions: []dynamodb.Table_AttributeDefinition{
			{
				AttributeName: "Username",
				AttributeType: "S",
			},
		},
		BillingMode: "PAY_PER_REQUEST",
		KeySchema: []dynamodb.Table_KeySchema{
			{
				AttributeName: "Username",
				KeyType:       "HASH",
			},
		},
		TableClass: "STANDARD_INFREQUENT_ACCESS",
		TableName:  "UserTable",
		Tags: []tags.Tag{
			{
				Key:   "Name",
				Value: "Username",
			},
		},
		AWSCloudFormationMetadata: map[string]interface{}{
			"Generator": "goformation",
		},
	}

	return template, nil
}
