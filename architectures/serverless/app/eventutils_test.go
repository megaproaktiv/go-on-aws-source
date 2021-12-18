package dsl_test

import (
	"fmt"
	"os"
	"testing"
    "io/ioutil"
	"encoding/json"
	"gotest.tools/assert"
	
	"github.com/aws/aws-lambda-go/events"
	"dsl"
)

func TestAppExtractObject(t *testing.T){
	var s3event events.S3Event;

	const testfile = "test/put.json"
	jsonFile, err := os.Open(testfile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully Opened ", testfile)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		print(err)
	}

	err = json.Unmarshal([]byte(byteValue), &s3event)


	assert.Equal(t,nil,err);

	expectedKey := "my2etestkey.txt"
	realKey := dsl.ExtractKey(s3event);

	assert.Equal(t, expectedKey,realKey)
	
}