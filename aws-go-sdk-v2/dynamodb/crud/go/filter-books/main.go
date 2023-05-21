package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var client *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = dynamodb.NewFromConfig(cfg)

}

func main() {

	tableName := "books"
	
	//begin filter
	params := &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		FilterExpression:          aws.String("begins_with(#EAISBN, :EAISBN)"),
		ExpressionAttributeNames: map[string]string{
			"#EAISBN" : "ISBN",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":EAISBN": &types.AttributeValueMemberS{
				Value: "111",
			},
		},
	}
	//end filter
	response, err := client.Scan(context.TODO(), params)
	if err != nil {
		log.Fatal("Error ddb get:", err)
	}

	fmt.Printf("Returned %v  items\n", len(response.Items))

	

}
