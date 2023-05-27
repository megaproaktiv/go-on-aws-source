package awsmockdemo_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"awsmockdemo"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	//begin import
	"github.com/megaproaktiv/awsmock"
	//end import
	"gotest.tools/assert"
)

func TestGetTableNameFile(t *testing.T) {
	//begin mockfunction
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
	//end mockfunction

	//begin test
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(GetParameterFunc)
	client := ssm.NewFromConfig(mockCfg.AwsConfig())
	name := awsmockdemo.GetTableName(client)
	
	assert.Equal(t, "totalfancyname",*name)
	//end test

}

func TestGetTableNameStruct(t *testing.T) {
	GetParameterFunc := func(ctx context.Context, params *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {		
		out := &ssm.GetParameterOutput{
			Parameter:      &types.Parameter{				
				Value:            aws.String("anothertotalfancyname"),
			},
		}
		return out,nil
	}

	// Create a Mock Handler
	mockCfg := awsmock.NewAwsMockHandler()
	// add a function to the handler
	// Routing per paramater types
	mockCfg.AddHandler(GetParameterFunc)

	// Create mocking client
	client := ssm.NewFromConfig(mockCfg.AwsConfig())

	name := awsmockdemo.GetTableName(client)
	assert.Equal(t, "anothertotalfancyname",*name)

}
