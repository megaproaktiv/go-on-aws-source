package kendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type KendraStackProps struct {
	awscdk.StackProps
}

func NewKendraStack(scope constructs.Construct, id string, props *KendraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	this = stack

	kendraPolicy = awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
		AssignSids: aws.Bool(true),
		Statements: &[]awsiam.PolicyStatement{
			&awsiam.NewPolicyStatement(
				&awsiam.PolicyStatementProps{
					Actions:       &[]*string{ 
						aws.String("cloudwatch:PutMetricData"),
					},
					Conditions:    &map[string]interface{}{
						"StringEquals": {
							"cloudwatch:namespace": "AWS/Kendra"
						}
					},
					Effect:        awsiam.Effect_ALLOW
				},
			),
			&awsiam.NewPolicyStatement(
				&awsiam.PolicyStatementProps{
					Actions:       &[]*string{
						"logs:DescribeLogGroups"
					},
					Resources:     &[]*string{
						aws.String("*")
					},
				}
			)

			},
		},
	})

	return stack
}
