package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func main() {
  //begin
	var event events.S3Event
	data, err := os.ReadFile("s3event.json")
	if err != nil {
		fmt.Println("Cant read input testdata: ", err)
	}
	err = json.Unmarshal(data, &event)
	if err != nil {
		fmt.Println("Cant unmarshal input testdata: ", err)
	}
	fmt.Println("Bucket: ",event.Records[0].S3.Bucket.Name)
	fmt.Println("Key: ",event.Records[0].S3.Object.Key)
	//end
}
