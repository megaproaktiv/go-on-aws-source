package eventutils_test

import (
	// Prepare Testdata
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
	
	// The function to be tested
	"eventutils"
	"github.com/aws/aws-lambda-go/events"
	
	// Imports for automated testing
	"testing"
	"gotest.tools/assert"
)

func TestExtractKey(t *testing.T) {
	var s3event events.S3Event

	const testfile = "testdata/put.json"

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

	assert.Equal(t, nil, err)

	expectedKey := "object-key-demo3"
	realKey := eventutils.ExtractKey(s3event)
	assert.Equal(t, expectedKey, realKey)

}
