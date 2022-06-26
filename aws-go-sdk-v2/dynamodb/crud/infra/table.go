package table

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-sdk-go-v2/aws"


	// "github.com/aws/jsii-runtime-go"
)

type TableStackProps struct {
	awscdk.StackProps
}

func NewTableStack(scope constructs.Construct, id string, props *TableStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	this := stack

	awsdynamodb.NewTable(this, aws.String("crud"), &awsdynamodb.TableProps{
		PartitionKey:               &awsdynamodb.Attribute{
				Name: aws.String("ID"),
				Type: awsdynamodb.AttributeType_STRING,
		},
		BillingMode:                awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TableClass:                 awsdynamodb.TableClass_STANDARD,
		TableName:                  aws.String("crud"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,

	})

	return stack
}
