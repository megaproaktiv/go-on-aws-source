package dsl_test

import (
	"context"
	"log"
	"os"
	"time"

	"testing"

	paddle "github.com/PaddleHQ/go-aws-ssm"
	"gotest.tools/assert"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	
)

var params *paddle.Parameters
var ClientD *dynamodb.Client
var ClientS *s3.Client

func init() {
	if os.Getenv("I_TEST") == "yes" {
		pmstore, err := paddle.NewParameterStore()
		if err != nil {
			log.Fatal("Cant connect to Parameter Store")
		}
		//Requesting the base path
		params, err = pmstore.GetAllParametersByPath("/goa-serverless/", true)
		if err != nil {
			log.Fatal("Cant get Parameter Store")
		}

		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("unable to load SDK config, " + err.Error())
		}
		// Using the Config value, create the DynamoDB client
		ClientD = dynamodb.NewFromConfig(cfg)
		ClientS = s3.NewFromConfig(cfg)
	}

}

// Integration Test
// export I_TEST="yes"
func TestAppInvokeLambdaWithEvent(t *testing.T) {
	if os.Getenv("I_TEST") != "yes" {
		t.Skip("Skipping testing in non Integration environment")
	  }
	table := params.GetValueByName("table")
	bucket := params.GetValueByName("bucket")

	var itemID = "my2etestkey.txt"

	key := map[string]types.AttributeValue{
		"itemID": &types.AttributeValueMemberS{Value: itemID},
	}
	// *** Delete item from table ****
	// Unless you specify conditions, the DeleteItem is an idempotent operation;
	// running it multiple times on the same item or attribute does not result in an error response.
	parmsDDBDelete := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: &table,
	}
	t.Log("Setup - delete item")
	_, err := ClientD.DeleteItem(context.TODO(), parmsDDBDelete)
	assert.NilError(t, err, "Delete Item should work")
	parmsDDBGet := &dynamodb.GetItemInput{
		Key:                      key,
		TableName:                &table,
		ConsistentRead:           aws.Bool(true),
	}
	responseDDB, err := ClientD.GetItem(context.TODO(), parmsDDBGet)
	assert.Equal(t, 0, len(responseDDB.Item), "Delete should work")

	// *** Copy object to S3
	testObjectFilename := "./testdata/dummy.txt"
	file, err := os.Open(testObjectFilename)
	assert.NilError(t, err, "Open file "+testObjectFilename+" should work")

	defer file.Close()
	t.Log("Copy object to S3")
	parmsS3 := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &itemID,
		Body:   file,
	}
	_, err = ClientS.PutObject(context.TODO(), parmsS3)
	assert.NilError(t, err, "Put Object "+itemID+" should work")

	t.Log("Sleep 3 seconds")
	time.Sleep(3 * time.Second)

	t.Log("Test item on DynamoDB")
	responseDDB, err = ClientD.GetItem(context.TODO(), parmsDDBGet)
	assert.NilError(t, err, "Get Item should work")
	// Item itsef with attribute time
	assert.Equal(t, 2, len(responseDDB.Item))

}

