package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client *dynamodb.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = dynamodb.NewFromConfig(cfg)

}

type LanguageStatus struct {
	ID string       `dynamodbav:"ID"`
	Status  string	`dynamodbav:"Status"`
}



func main() {

	languageID := "GOCLIENT2"

	languageStatusItem := &LanguageStatus{
		ID:     languageID,
		Status: "MORETHANOK",
	}

	ddbItem, err := attributevalue.MarshalMap(languageStatusItem)
	if err != nil {
		log.Fatal("failed to marshal Record, %w", err)
	}

	params := &dynamodb.PutItemInput{
		Item: ddbItem,
		TableName:                   aws.String("crud"),
	}

	_, err = client.PutItem(context.TODO(), params)
	if err != nil {
		log.Fatal("Error ddb put:", err)
	}

}

