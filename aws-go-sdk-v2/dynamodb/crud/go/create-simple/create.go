package create

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var Client *dynamodb.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = dynamodb.NewFromConfig(cfg)

}

type LanguageStatus struct {
	ID string       `dynamodbav:"ID"`
	Status  string	`dynamodbav:"Status"`
}

type LanguagesField string

const (
	LanguageStatus_FieldID      LanguagesField = "ID"
	LanguageStatus_FieldStatus  LanguagesField = "Status"
)

func Create(client *dynamodb.Client) {

	languageID := "GOCLIENT"

	params := &dynamodb.PutItemInput{
		Item:  map[string]types.AttributeValue{
			string(LanguageStatus_FieldID): &types.AttributeValueMemberS{
				Value:  languageID,
			},
			string(LanguageStatus_FieldStatus): &types.AttributeValueMemberS{
				Value:  "OK",
			},
		},
		TableName:                   aws.String("crud"),
	}

	_, err := client.PutItem(context.TODO(), params)
	if err != nil {
		log.Fatal("Error with put:", err)
	}

}

