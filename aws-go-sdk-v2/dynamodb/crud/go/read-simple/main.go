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

	//begin get
	tableName := "barjokes"
	NAME := "moebius"
	params := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			string("NAME"): &types.AttributeValueMemberS{
				Value: NAME,
			},
		},
		TableName: aws.String(tableName),
	}

	response, err := client.GetItem(context.TODO(), params)
	//end get

	if err != nil {
		log.Fatal("Error ddb get:", err)
	}

	//begin read
	for k, item := range response.Item {
		var content string
		// type switches can be used to check the union value
		switch v := item.(type) {
			case *types.AttributeValueMemberN:
				content = v.Value // Value is number
			case *types.AttributeValueMemberS:
				content = v.Value // Value is string
			default:
				fmt.Println("nil or unknown type")
		}
		fmt.Printf("%s: %v \n", k, content)
	}
	//end read

}
