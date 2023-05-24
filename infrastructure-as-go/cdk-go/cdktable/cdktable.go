// begin package
package cdktable

//end package

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"

	//begin import
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	//end import
)

// begin props
type CdktableStackProps struct {
	StackProps awscdk.StackProps
}

// end props
func NewCdktableStack(scope constructs.Construct, id string, props *CdktableStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	//begin this
	stack := awscdk.NewStack(scope, &id, &sprops)
	this := stack
	//end this

	//begin table
	table := dynamodb.NewTable(this, aws.String("barjokes"), &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: aws.String("NAME"),
			Type: dynamodb.AttributeType_STRING,
		},
		BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   aws.String("barjokes-cdk"),
	})
	//end table

	//begin output
	awscdk.NewCfnOutput(this, aws.String("tablename"), &awscdk.CfnOutputProps{
		Value: table.TableName(),
	})
	//end output

	return stack
}
