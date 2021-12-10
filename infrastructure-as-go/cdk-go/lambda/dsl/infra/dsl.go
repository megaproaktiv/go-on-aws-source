package main

import (
	"log"
	"os"
	"path/filepath"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	// "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	
	
	logs "github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"	
)

type DslStackProps struct {
	awscdk.StackProps
}

func NewDslStack(scope constructs.Construct, id string, props *DslStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lambdaPath := filepath.Join(path, "../dist/main.zip")

	myHandler := awslambda.NewFunction(stack, aws.String("myHandler"), 
	&awslambda.FunctionProps{
		Description:                  aws.String("dsl - dynamodb s3 lambda"),
		FunctionName:                 aws.String("logincomingobject"),
		LogRetention:                 logs.RetentionDays_THREE_MONTHS,
		MemorySize:                   aws.Float64(1024),
		Timeout:                      awscdk.Duration_Seconds(aws.Float64(10)),
		Code: awslambda.Code_FromAsset(&lambdaPath, &awss3assets.AssetOptions{}),
		Handler: aws.String("main"),
		Runtime: awslambda.Runtime_GO_1_X(),
	})

	// Bucket start ****************
    // *
    bucky := awss3.NewBucket(stack, aws.String("incoming-gov2"), &awss3.BucketProps{
    	BlockPublicAccess:      awss3.BlockPublicAccess_BLOCK_ALL(),
    });
	
	// Tell Lambda the dynamic bucket name
	myHandler.AddEnvironment(aws.String("Bucket"), bucky.BucketName(), nil);
	// *
	// give lambda read rights
	bucky.GrantRead(myHandler, aws.String("*"))
	// *
	// Bucket end *******************

	//** Dynamodb start */
    // do not force table name, this leads to singleTimeDeployability
    myTable := awsdynamodb.NewTable(stack, aws.String("items"), &awsdynamodb.TableProps{
    	PartitionKey:               &awsdynamodb.Attribute{
    		Name: aws.String("itemID"),
    		Type: awsdynamodb.AttributeType_STRING,
    	},
    	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
    })
  
	myTable.GrantReadWriteData(myHandler);
	myHandler.AddEnvironment(aws.String("TableName"), myTable.TableName(), nil);


	return stack
}
