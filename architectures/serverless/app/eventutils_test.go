package trick_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-lambda-go/events"
    "io/ioutil"
	"os"
	"testing"
	"trick"
)

func TestExtractObject(t *testing.T){
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

	expectedKey := "object-key"
	realKey := trick.ExtractKey(s3event);

	assert.Equal(t, expectedKey,realKey)
	
}