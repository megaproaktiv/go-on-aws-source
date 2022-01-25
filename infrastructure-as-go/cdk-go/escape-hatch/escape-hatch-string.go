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
// Show 
// - not working direct string assigment
// - String as interface{}
func NewEscapeHatchStringStack(scope constructs.Construct, id string, props *EscapeHatchStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	helper := awss3.NewBucket(stack, aws.String("helper"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	bucketStruct := awss3.NewBucket(stack, aws.String("bucky"), &awss3.BucketProps{
		BlockPublicAccess: awss3.BlockPublicAccess_BLOCK_ALL(),
	})

	var cfnBucketStruct awss3.CfnBucket

	jsii.Get(bucketStruct.Node(), "defaultChild", &cfnBucketStruct)

	// Direct string assigment does not work
	cfnBucketStruct.AddPropertyOverride(aws.String("AnalyticsConfigurationsEscaped"), `
	[
    {
    "Id": "AnalyticsConfigurationId",
    "StorageClassAnalysis": {
        "DataExport": {
            "Destination": {
                "BucketArn": {
                    "Fn::GetAtt": [
                        "Helper",
                        "Arn"
                    ]
                },
                "Format": "CSV",
                "Prefix": "AnalyticsDestinationPrefix"
            },
            "OutputSchemaVersion": "V_1"
        }
    },
    "Prefix": "AnalyticsConfigurationPrefix",
    "TagFilters": [
        {
            "Key": "AnalyticsTagKey",
            "Value": "AnalyticsTagValue"
        }
    ]
    }
]
`)

// This will generate the wrong Cfn Code.
// The Analytics Configuration has to be an array
cfnBucketStruct.AddPropertyOverride(aws.String("AnalyticsConfigurations"),
	&map[string]interface{}{
	
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
)

	return stack
}
