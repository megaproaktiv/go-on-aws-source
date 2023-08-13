package dsl

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	event "github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	logs "github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"	
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"	
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

	// Lambda start ***********
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	//begin lambda
	lambdaPath := filepath.Join(path, "../app/dist/main.zip")
	
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
	//end lambda
	awsssm.NewStringParameter(stack, aws.String("handler-output"), 
		&awsssm.StringParameterProps{
			Description:    aws.String("Store lambda function name"),
			ParameterName:  aws.String("/goa-serverless/handler"),
			StringValue: myHandler.FunctionName(),
		},
	)
	// Lambda end ***********

	// Bucket start ****************
    // *
	//begin bucket
    bucky := awss3.NewBucket(stack, aws.String("incoming-gov2"), &awss3.BucketProps{
    	BlockPublicAccess:      awss3.BlockPublicAccess_BLOCK_ALL(),
    });	
	//end bucket
	awsssm.NewStringParameter(stack, aws.String("bucket-output"), 
	&awsssm.StringParameterProps{
		Description:    aws.String("Store bucket name"),
		ParameterName:  aws.String("/goa-serverless/bucket"),
		StringValue: bucky.BucketName(),
	},
	)
	// Tell Lambda the dynamic bucket name
	//begin instrumentation
	myHandler.AddEnvironment(aws.String("Bucket"), bucky.BucketName(), nil);
	// give lambda read rights
	bucky.GrantRead(myHandler, aws.String("*"))
	//end instrumentation
	// *
	// Bucket end *******************

	//begin event  *******************
	myHandler.AddEventSource(event.NewS3EventSource(bucky, &event.S3EventSourceProps{
		Events:  &[]awss3.EventType{awss3.EventType_OBJECT_CREATED,},
	}))
	//end event  *******************

	//** Dynamodb start ******
    // do not force table name, this leads to singleTimeDeployability
	//begin table
    myTable := awsdynamodb.NewTable(stack, aws.String("items"), &awsdynamodb.TableProps{
    	PartitionKey:               &awsdynamodb.Attribute{
    		Name: aws.String("itemID"),
    		Type: awsdynamodb.AttributeType_STRING,
    	},
    	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
    })
  
	myTable.GrantReadWriteData(myHandler);
	myHandler.AddEnvironment(aws.String("TableName"), myTable.TableName(), nil);
	//end table
	awsssm.NewStringParameter(stack, aws.String("table-output"), 
		&awsssm.StringParameterProps{
			Description:    aws.String("Store ltable name"),
			ParameterName:  aws.String("/goa-serverless/table"),
			StringValue: myTable.TableName(),
		},
	)
	//** Dynamodb end ******
	
	return stack
}
