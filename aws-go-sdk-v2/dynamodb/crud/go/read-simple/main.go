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

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = dynamodb.NewFromConfig(cfg)

}




func main() {

	languageID := "GOCLIENT"

	params := &dynamodb.GetItemInput{
		Key:                      map[string]types.AttributeValue{
			string("ID"): &types.AttributeValueMemberS{
				Value:  languageID,
			},
		},
		TableName:                   aws.String("crud"),

	}

	response, err := client.GetItem(context.TODO(), params)
	if err != nil {
		log.Fatal("Error ddb put:", err)
	}
	for i, item := range response.Item{
		var content string
		// type switches can be used to check the union value
		switch v := item.(type) {
		case *types.AttributeValueMemberN:
			content = v.Value // Value is string
		case *types.AttributeValueMemberS:
			content = v.Value // Value is string
		case *types.AttributeValueMemberSS:
			_ = v.Value // Value is []string
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		default:
			fmt.Println("nil or unknown type")

		}
		fmt.Printf("Response number %v: %v \n",i,content)
	}

}


