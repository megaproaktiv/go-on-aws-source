package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-sdk-go-v2/aws"	
)

type DynamodbCdkGoStackProps struct {
	awscdk.StackProps
}

func NewDynamodbCdkGoStack(scope constructs.Construct, id string, props *DynamodbCdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	this := stack
	table := dynamodb.NewTable(this, aws.String("table"), &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: aws.String("username"),
			Type: dynamodb.AttributeType_STRING,
		},
		BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
		TableName: aws.String("Username"),
	})

	awscdk.Tags_Of(table).Add(aws.String("Name"), aws.String("Username"),nil)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewDynamodbCdkGoStack(app, "DynamodbCdkGoStack", &DynamodbCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
