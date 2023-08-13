// begin main
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hiho %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}

//end main
