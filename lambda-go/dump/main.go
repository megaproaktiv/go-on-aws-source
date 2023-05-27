package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)
//begin dump
func HandleRequest(ctx context.Context, event interface{}) (string, error) {
        eventJSON, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal event to JSON: %v", err)
                return "", err
	}
	fmt.Printf("Received event as JSON: %s\n", eventJSON)
        return string(eventJSON),nil
}
//end dump
func main() {
        lambda.Start(HandleRequest)
}