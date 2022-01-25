package escapehatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/awslabs/goformation/v5/cloudformation/s3"
)

func NewEscapeHatchStructStack(scope constructs.Construct, id string, props *EscapeHatchStackProps) awscdk.Stack {
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
		BucketName:        aws.String("bucky"),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	})

	var cfnBucketStruct awss3.CfnBucket

	jsii.Get(bucketStruct.Node(), "defaultChild", &cfnBucketStruct)

	analyticsConfiguration :=  &[]s3.Bucket_AnalyticsConfiguration{
		{
			Id:                                   "AnalyticsConfigurationId",
			Prefix:                               "AnalyticsConfigurationPrefix",
			StorageClassAnalysis:                 &s3.Bucket_StorageClassAnalysis{
				DataExport:                           &s3.Bucket_DataExport{
					Destination:                          &s3.Bucket_Destination{
						BucketAccountId:                      id,
						BucketArn:                            *helper.BucketArn(),
						Format:                               "CSV",
						Prefix:                               "AnalyticsDestinationPrefix",
					},
					OutputSchemaVersion:                  "V_1",
				},
			},
			TagFilters: []s3.Bucket_TagFilter{{
				Key: "AnalyticsTagKey",
				Value: "AnalyticsTagValue",
				},
			},
		},
	}
	
	cfnBucketStruct.AddPropertyOverride(aws.String("AnalyticsConfigurations"), *analyticsConfiguration)

	return stack
}
