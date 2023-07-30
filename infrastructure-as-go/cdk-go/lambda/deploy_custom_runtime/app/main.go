package main

import (
	"context"
	"fmt"

	
	//begin import
	"github.com/aws/aws-lambda-go/lambda"
	//end import
)

// begin function
type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hiho %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}

//end function
