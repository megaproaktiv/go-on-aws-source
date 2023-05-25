package main


import (
	"context"
	"os"
	"dsl"
	//begin events
	"github.com/aws/aws-lambda-go/events"

	//end events
	"github.com/aws/aws-lambda-go/lambda"


)
// HandleRequest S3 Event
//begin events
func handler(ctx context.Context, s3Event events.S3Event) {
//end events
	// See https://github.com/aws/aws-lambda-go/tree/master/events
	// Handle only one event
	//begin logic
	s3input := dsl.ExtractKey(s3Event);
	tableName := os.Getenv("TableName")

	err := dsl.PutItem(dsl.Client, s3input,tableName)
	//end logic
	if err != nil{
		panic(err)
	}
}


func main() {

	lambda.Start(handler)

}

