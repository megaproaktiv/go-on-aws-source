package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"trick"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go"
)
// Hier geht dann langsam die Übersicht verloren
// HandleRequest S3 Event
func handler(ctx context.Context, s3Event events.S3Event) {
	// See https://github.com/aws/aws-lambda-go/tree/master/events
	// Handle only one event
	s3input := trick.ExtractKey(s3Event);
	tableName := os.Getenv("TableName")

	putItem(s3input,tableName)

}

func putItem(itemID string, tableName string){
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(config.WithRegion("eu-west-1"))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}



	// Set the AWS Region that the service clients should use
	cfg.Region = "eu-west-1"

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	t := time.Now()


	input := &dynamodb.PutItemInput{
        Item: map[string]*types.AttributeValue{
            "itemID": {
                S: aws.String(itemID),
			},
			"time" : {
				S: aws.String(t.String()),
			},
        },
        TableName: aws.String(tableName),
    }

	// Build the request with its input parameters
	result, err := client.PutItem(context.TODO(),input)
	if err != nil {
    // To get a specific API error
		var notFoundErr *types.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
			log.Printf("scan failed because the table was not found, %v",
				notFoundErr.ErrorMessage())
		}

		// To get any API error
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("scan failed because of an API error, Code: %v, Message: %v",
				apiErr.ErrorCode(), apiErr.ErrorMessage())
		}

		// To get the AWS response metadata, such as RequestID
		var respErr *awshttp.ResponseError // Using import alias "awshttp" for package github.com/aws/aws-sdk-go-v2/aws/transport/http
		if errors.As(err, &respErr) {
			log.Printf("scan failed with HTTP status code %v, Request ID %v and error %v",
				respErr.HTTPStatusCode(), respErr.ServiceRequestID(), respErr)
		}

		
	}


	fmt.Println("Response", result)
}

func main() {

	lambda.Start(handler)

}

