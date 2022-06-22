package create_test

import (
	"context"
	"create"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"

	"github.com/megaproaktiv/awsmock"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	
	PutItemFunc := func(ctx context.Context, params *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {		
		var content string

		var err error = errors.New("nil or unknown type")
		switch v := params.Item["ID"].(type) {
		case *types.AttributeValueMemberN:
			content = v.Value // Value is string
			err = errors.New("Should be S member")
		case *types.AttributeValueMemberS:
			content = v.Value // Value is strin
			err = nil
			assert.Equal(t, content, "GOCLIENT")
		default:
			assert.Error(t,err,"nil or unknown type")
		}

		out := &dynamodb.PutItemOutput{
		}
		return out,nil
	}

	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(PutItemFunc)
	client := dynamodb.NewFromConfig(mockCfg.AwsConfig())

	create.Create(client)
	

}
