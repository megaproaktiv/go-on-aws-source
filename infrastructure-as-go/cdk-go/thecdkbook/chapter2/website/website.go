package website

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	origins "github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"

	// SDK
	aws "github.com/aws/aws-sdk-go-v2/aws"
)

type WebsiteStackProps struct {
	awscdk.StackProps
}

func NewWebsiteStack(scope constructs.Construct, id string, props *WebsiteStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	this := awscdk.NewStack(scope, &id, &sprops)


	myBucket := awss3.NewBucket(this, aws.String("MyBucket"),nil)

	awscloudfront.NewDistribution(this, aws.String("MyDistribution"),
			&awscloudfront.DistributionProps{
				DefaultBehavior: &awscloudfront.BehaviorOptions{
					Origin: origins.NewS3Origin(myBucket, nil),
				},
			},
		)

	return this
}

