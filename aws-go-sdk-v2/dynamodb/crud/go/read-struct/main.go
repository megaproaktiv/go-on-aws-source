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

type LanguageStatus struct {
	ID string       `dynamodbav:"ID"`
	Status  string	`dynamodbav:"Status"`
}



func main() {

	languageID := "GOCLIENT2"

	

	// AttributeValue Unmarshaling 
	// To unmarshal an AttributeValue to a Go type you can use the Unmarshal,
	//  UnmarshalList, UnmarshalMap, and UnmarshalListOfMaps functions.
	// The List and Map functions are specialized versions of the Unmarshal function 
	// for unmarshal slices and maps of Attributevalues.
	params := &dynamodb.GetItemInput{
		Key:                      map[string]types.AttributeValue{
			string("ID"): &types.AttributeValueMemberS{
				Value:  languageID,
			},
		},
		TableName:                   aws.String("crud"),
	}

	result, err := client.GetItem(context.TODO(), params)
	if err != nil {
		log.Fatal("Error ddb get:", err)
	}

	// type GetItemOutput 
	// has 
	// Item map[string]types.AttributeValue
	// as variable, so use UnmarshalMap

	var item LanguageStatus
	err = attributevalue.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Fatal("failed to unmarshal Items, %w", err)
	}

	// Now item struct is filled
	fmt.Printf("Status is: %v\n",item.Status)
	
}

