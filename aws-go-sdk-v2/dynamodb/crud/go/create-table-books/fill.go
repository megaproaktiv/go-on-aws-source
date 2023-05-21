package table

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Books struct {
	Id     string `json:"Id" dynamodbav:"Id"`
	ISBN string `json:"ISBN"`
	Title  string `json:"Title"`
}

func FillTable(table *string) error{

	items,err := ReadFile();
	if err != nil {	
		return err
	}


	// Put each item into DynamoDB
	for _, record := range items {
		av, err := attributevalue.MarshalMap(record)
		if err != nil {
			fmt.Println("failed to marshal item:", err)
			return err
		}

		input := &dynamodb.PutItemInput{
			TableName: table,
			Item:      av,
		}

		_ ,err = Client.PutItem(context.TODO(), input)
		if err != nil {
			fmt.Println("failed to write item to DynamoDB:", err)
			return err
		}
		
	}
	return nil
}

func ReadFile() ([]Books, error) {
		// Read JSON file
		fileBytes, err := os.ReadFile("items.json")
		if err != nil {
			fmt.Println("failed to read JSON file:", err)
			return nil,err
		}
	
		// Unmarshal JSON into a slice of items
		var items []Books
		err = json.Unmarshal(fileBytes, &items)
		if err != nil {
			fmt.Println("failed to unmarshal JSON:", err)
		}
		return items,nil
}