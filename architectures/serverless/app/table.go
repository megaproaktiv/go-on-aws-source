package dsl

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/smithy-go"
)

var Client *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	// Using the Config value, create the DynamoDB client
	Client = dynamodb.NewFromConfig(cfg)
}

func PutItem(client *dynamodb.Client, itemID string, tableName string) error {

	t := time.Now()
	//begin dynamodbput
	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"itemID": &types.AttributeValueMemberS{
				Value: itemID,
			},
			"time": &types.AttributeValueMemberS{
				Value: t.String(),
			},
		},
		TableName: aws.String(tableName),
	}
	//end dynamodbput

	// Build the request with its input parameters
	_, err := client.PutItem(context.TODO(), input)
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

		return err

	}

	return nil
}
