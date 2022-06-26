package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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
	
	// Filter
	filter := expression.Name("ID").BeginsWith("GO")
	
	// Build the dynamodb syntax for querys
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("crud"),
	}

	result, err := client.Scan(context.TODO(), params)
	if err != nil {
		log.Fatal("Error ddb get:", err)
	}

	// type ScanOutput  
	// has 
	// Items []map[string]types.AttributeValue
	// as variable, so use UnmarshalMap

	items := []LanguageStatus{}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}


	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		log.Fatal("failed to unmarshal Items, %w", err)
	}

	// Now items struct is filled
	for _,item := range items {
		fmt.Printf("Status is: %v\n",item.Status)
	}
	
}

