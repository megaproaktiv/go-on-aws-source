package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var client *dynamodb.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = dynamodb.NewFromConfig(cfg)

}

//begin struct
type BarJoke struct {
	Name     string `dynamodbav:"NAME"`
	Rating   float64    `dynamodbav:"rating"`
	Headline string `dynamodbav:"headline"`
	Content  string `dynamodbav:"content"`
}
//end struct


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
	//begin unmarshal
	var joke BarJoke
	attributevalue.UnmarshalMap(response.Item,&joke)
	//end unmarshal
	if err != nil {
		log.Fatal("Error UnmarshalMap :", err)
	}

	//begin output
	fmt.Printf("Headline: %v\n", joke.Headline)
	//end output
}

