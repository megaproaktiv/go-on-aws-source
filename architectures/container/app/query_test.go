package showtable_test

import (
	"context"
	"showtable"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/megaproaktiv/awsmock"
	"gotest.tools/assert"
)

func TestQueryDDB(t *testing.T) {
	ScanFunc := func(ctx context.Context, params *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {	
		
		out := &dynamodb.ScanOutput{
			ConsumedCapacity: &types.ConsumedCapacity{},
			Count:            2,
			Items:            []map[string]types.AttributeValue{
				{
					"itemID" : &types.AttributeValueMemberS{
						Value: "object-key-demo1",
					},
				},
				{
					"itemID" : &types.AttributeValueMemberS{
						Value: "object-key-demo2",
					},
				},
			},
			LastEvaluatedKey: map[string]types.AttributeValue{},
			ScannedCount:     2,
			ResultMetadata:   middleware.Metadata{},
		}
		
		return out,nil
	}

	// Create a Mock Handler
	mockCfg := awsmock.NewAwsMockHandler()
	// add a function to the handler
	// Routing per paramater types
	mockCfg.AddHandler(ScanFunc)

	client := dynamodb.NewFromConfig(mockCfg.AwsConfig())

	results := showtable.QueryDDB(client, aws.String("mytable"))
	
	assert.Equal(t, "object-key-demo1",results[0].ItemID)
	assert.Equal(t, "object-key-demo2",results[1].ItemID)
}
