package table

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type TableProps struct {
	cdk.StackProps
}

func TableStack(scope constructs.Construct, id string, props *TableProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)
	table := awsdynamodb.NewTable(stack, aws.String("table"),
		&awsdynamodb.TableProps{
			PartitionKey: &awsdynamodb.Attribute{
				Name: aws.String("itemID"),
				Type: awsdynamodb.AttributeType_STRING,
			},
			// NOT recommended for production code
			RemovalPolicy: cdk.RemovalPolicy_DESTROY,
		})

	tableParm := "/go-on-aws/table"
	awsssm.NewStringParameter(stack, &tableParm,
		&awsssm.StringParameterProps{
			AllowedPattern: new(string),
			Description:    aws.String("dynamodb Table name go on aws architecture"),
			ParameterName:  &tableParm,
			StringValue:    table.TableName(),
		})

	return stack
}
