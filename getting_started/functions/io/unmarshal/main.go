package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/events"
)

func main() {
	var event events.S3Event
	data, err := os.ReadFile("s3event.json")
	if err != nil {
		fmt.Println("Cant read input testdata: ", err)
	}
	json.Unmarshal(data, &event)

	fmt.Println(event.Records[0].S3.Bucket.Name)
	fmt.Println(event.Records[0].S3.Object.Key)
}
