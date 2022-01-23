package escapehatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EscapeHatchStackProps struct {
	awscdk.StackProps
}

func NewEscapeHatchStringStack(scope constructs.Construct, id string, props *EscapeHatchStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	helper := awss3.NewBucket(stack, aws.String("helper"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// example resource
	bucketStruct := awss3.NewBucket(stack, aws.String("bucky"), &awss3.BucketProps{
		BlockPublicAccess: awss3.BlockPublicAccess_BLOCK_ALL(),
	})

	var cfnBucketStruct awss3.CfnBucket

	jsii.Get(bucketStruct.Node(), "defaultChild", &cfnBucketStruct)
	// Example from https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#aws-properties-s3-bucket-analyticsconfiguration--examples

	// have to use json as struct
	// if you use json as string, all " will be escaped to \"
	// with structures it will render ok
	cfnBucketStruct.AddPropertyOverride(aws.String("AnalyticsConfigurations"),
		&[]map[string]interface{}{
			{
				"Id": "AnalyticsConfigurationId",
				"StorageClassAnalysis": map[string]interface{}{
					"DataExport": map[string]interface{}{
						"Destination": map[string]interface{}{
							"BucketArn": helper.BucketArn(),
							"Format": "CSV",
							"Prefix": "AnalyticsDestinationPrefix",
						},
						"OutputSchemaVersion": "V_1",
					},
				},
			},
		},
	)

	return stack
}
