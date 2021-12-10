package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// See
// https://github.com/aws/aws-lambda-go/blob/main/cfn/README.md
// for using a CustomResource Wrapper


func HandleRequest(ctx context.Context, event cfn.Event) ( interface{}, error)  {
	
	out, err := json.MarshalIndent(event, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	r := cfn.NewResponse(&event)
	if r.PhysicalResourceID == "" {
		log.Println("PhysicalResourceID must exist on creation, copying Log Stream name")
		r.PhysicalResourceID = lambdacontext.LogStreamName
	}
	
	switch event.RequestType  {
		case cfn.RequestCreate:  
			r.Status = cfn.StatusSuccess
		case cfn.RequestUpdate:
			r.Status = cfn.StatusSuccess
		case cfn.RequestDelete:
			r.Status = cfn.StatusSuccess
		default:
			fmt.Println("The value is wrong")	
	}
	
	err = r.Send()
		if err != nil {
			log.Fatal(err)
		}
	
	return "Status: ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}

