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
	
	//begin query
	params := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:          aws.String("#EAID = :EAID"),
		ExpressionAttributeNames: map[string]string{
			"#EAID" : "Id",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":EAID": &types.AttributeValueMemberS{
				Value: "101",
			},
		},
	}
	//end query
	response, err := client.Query(context.TODO(), params)
	if err != nil {
		log.Fatal("Error ddb get:", err)
	}

	fmt.Printf("Returned %v  items\n", len(response.Items))



}
