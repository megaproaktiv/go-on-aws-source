package escapehatch

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/awslabs/goformation/v5/cloudformation/s3"
)

func NewEscapeHatchFileStack(scope constructs.Construct, id string, props *EscapeHatchStackProps) awscdk.Stack {
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
	bucky := awss3.NewBucket(stack, aws.String("bucky"), &awss3.BucketProps{
		BlockPublicAccess: awss3.BlockPublicAccess_BLOCK_ALL(),
	})

	var cfnBucketStruct awss3.CfnBucket

	jsii.Get(bucky.Node(), "defaultChild", &cfnBucketStruct)

	var analyticsConfigurationFromFile []s3.Bucket_AnalyticsConfiguration

	data, err := os.ReadFile("testdata/analyticsconfig.json")
	if err != nil {
		fmt.Println("Cant read json data: ", err)
	}
	json.Unmarshal(data, &analyticsConfigurationFromFile)
	if err != nil {
		fmt.Println("JSON unmarshall error data: ", err)
	}
	analyticsConfigurationFromFile[0].StorageClassAnalysis.DataExport.Destination.BucketArn = *helper.BucketArn()
	cfnBucketStruct.AddPropertyOverride(aws.String("AnalyticsConfigurations"), analyticsConfigurationFromFile)

	return stack
}
