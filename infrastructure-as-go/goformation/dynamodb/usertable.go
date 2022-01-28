package phantom

import (
	"fmt"

	"github.com/awslabs/goformation/v5/cloudformation"
	"github.com/awslabs/goformation/v5/cloudformation/dynamodb"
	"github.com/awslabs/goformation/v5/cloudformation/tags"
)

// Synthesize CloudFormation
// for a simple Dynamodb User Table
func Synth() (*string, error) {

	template := cloudformation.NewTemplate()
	template.Resources["simpletable"] = &dynamodb.Table{
		AttributeDefinitions:                 []dynamodb.Table_AttributeDefinition{
			{
				AttributeName: "Username",
				AttributeType: "S",
			},
		},
		BillingMode:                          "PAY_PER_REQUEST",
		KeySchema:                            []dynamodb.Table_KeySchema{
			{
				AttributeName:                        "Username",
				KeyType:                              "HASH",
			},
		},
		TableClass:                           "STANDARD_INFREQUENT_ACCESS",
		TableName:                            "UserTable",
		Tags:                                 []tags.Tag{
			{
				Key: "Name",
				Value: "Username",
			},
		},
		AWSCloudFormationMetadata:            map[string]interface{}{
			"Generator" : "goformation",
		},
	}

	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
		return nil, err
	} 

	content := string(y)

	return &content, nil
}