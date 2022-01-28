package showtable_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"showtable"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/megaproaktiv/awsmock"
	"gotest.tools/assert"

)

func TestGetTableName(t *testing.T) {
	GetParameterFunc := func(ctx context.Context, params *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {	
		
		testfile := "testdata/ssm-get-parameter.json"
		data, err := ioutil.ReadFile(testfile)
		if err != nil {
			fmt.Println("File reading error: ", err)
		}
		
		out := &ssm.GetParameterOutput{}
		err = json.Unmarshal(data, out); 
		if err != nil {
			t.Error(err)
		}
		return out,nil
	}

	// Create a Mock Handler
	mockCfg := awsmock.NewAwsMockHandler()
	// add a function to the handler
	// Routing per paramater types
	mockCfg.AddHandler(GetParameterFunc)

	client := ssm.NewFromConfig(mockCfg.AwsConfig())

	name := showtable.GetTableName(client)
	
	
	assert.Equal(t, "objectlister-items07D08F4B-KSHYJ5OP4BRW",*name)

}
