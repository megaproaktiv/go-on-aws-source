package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

//begin struct
type BarJoke struct {
	Name     string `dynamodbav:"NAME"`
	Rating   float64    `dynamodbav:"rating"`
	Headline string `dynamodbav:"headline"`
	Content  string `dynamodbav:"content"`
}
//end struct




func main() {
	tableName := "barjokes"
	//begin filter
	filter := expression.Name("Headline").Contains("bartender")

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	//end filter

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return
	}

	//begin params
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 &tableName,
	}

	result, err := client.Scan(context.TODO(), params)
	//end params
	if err != nil {
		log.Fatal("Error ddb get:", err)
	}

	// type ScanOutput
	// has
	// Items []map[string]types.AttributeValue
	// as variable, so use UnmarshalMap

	items := []BarJoke{}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}


	// Now items struct is filled
	for _,item := range items {
		fmt.Printf("Headline is: %v\n",item.Headline)
	}

}
