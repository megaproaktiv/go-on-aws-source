package dsl_test

import (
	"reflect"
	"testing"

	"github.com/megaproaktiv/awsmock"
	"gotest.tools/assert"

	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dsl"
)

const testkey = "kjhkjjhollymolly"


func TestPutItem(t *testing.T) {
	t.Log("App - DynamodD Put Item")
	PutItemMock := func(ctx context.Context, params *dynamodb.PutItemInput)(*dynamodb.PutItemOutput, error) {
		
		item := params.Item["itemID"]
		
		
		expected := &types.AttributeValueMemberS{
			Value: testkey,
		}
		assert.Equal(t,ptrToValue(expected),ptrToValue(item))
		return &dynamodb.PutItemOutput{}, nil
	}

	// Create a Mock Handler
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(PutItemMock)
	client := dynamodb.NewFromConfig(mockCfg.AwsConfig())



	err := dsl.PutItem(client,testkey, "egal")
	assert.NilError(t, err, "Error should be nil")

	// PutItem(tt.args.client, tt.args.itemID, tt.args.tableName); 

}

func ptrToValue(in interface{}) interface{} {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if !v.IsValid() {
		return nil
	}
	if v.Kind() == reflect.Ptr {
		return ptrToValue(v.Interface())
	}
	return v.Interface()
}