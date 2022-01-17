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

func NewEscapeHatchStack(scope constructs.Construct, id string, props *EscapeHatchStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	winterSoldier := awss3.NewBucket(stack, aws.String("bucky"), &awss3.BucketProps{
		BlockPublicAccess:      awss3.BlockPublicAccess_BLOCK_ALL(),
		BucketName:             aws.String("bucky"),
	})

	var cap awss3.CfnBucket

	jsii.Get(winterSoldier.Node(), "defaultChild", &cap)
	// Example from https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#aws-properties-s3-bucket-analyticsconfiguration--examples

	// have to use json as struct
	// if you use json as string, all " will be escaped to \"
	// with structures it will render ok
	cap.AddPropertyOverride(aws.String("AnalyticsConfigurations"),
		&map[string]interface{}{
			"Id": "AnalyticsConfigurationId",
			"StorageClassAnalysis": map[string]interface{}{
				"DataExport": map[string]interface{}{
					"Destination": map[string]interface{}{
						"BucketArn": winterSoldier.BucketArn(),
					},
					"Format": "CSV",
					"Prefix": "AnalyticsDestinationPrefix",
				},
				"OutputSchemaVersion": "V_1",
				"Prefix": "AnalyticsConfigurationPrefix",
				"TagFilters": []map[string]string{
                            {
                                "Key": "AnalyticsTagKey",
                                "Value": "AnalyticsTagValue",
                            },
				},
			},
		},
	)
	
	return stack
}
