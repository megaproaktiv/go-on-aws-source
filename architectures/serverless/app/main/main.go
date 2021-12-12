package main

import (
	"context"
	"os"
	"dsl"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"


)
// Hier geht dann langsam die Übersicht verloren
// HandleRequest S3 Event
func handler(ctx context.Context, s3Event events.S3Event) {
	// See https://github.com/aws/aws-lambda-go/tree/master/events
	// Handle only one event
	s3input := dsl.ExtractKey(s3Event);
	tableName := os.Getenv("TableName")

	err := dsl.PutItem(dsl.Client, s3input,tableName)
	if err != nil{
		panic(err)
	}
}


func main() {

	lambda.Start(handler)

}

