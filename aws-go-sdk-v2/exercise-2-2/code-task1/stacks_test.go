package crlist_test

import (
	"context"
	"crlist"
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"gotest.tools/assert"

	"github.com/megaproaktiv/awsmock"
)

func TestGetStatus(t *testing.T) {
	describeStackRessourcesMock := func(ctx context.Context, params *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error) {
		out := &cloudformation.DescribeStackResourcesOutput{}
		testfile := "testdata/describe_stack_resources.json"
		data, err := os.ReadFile(testfile)
		if err != nil {
			t.Errorf("cannnot read %v, error: %v", testfile, err)
			os.Exit(1)
		}
		err = json.Unmarshal(data, out)
		if err != nil {
			t.Errorf("cannot unmarshal %v, error: %v", testfile, err)
			os.Exit(2)
		}
		return out, nil
	}
	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(describeStackRessourcesMock)
	client := cloudformation.NewFromConfig(mockCfg.AwsConfig())
	result, err := crlist.GetStatus(client, aws.String("test"))
	assert.NilError(t, err)
	rArray := (*[2]crlist.ResourceStatus)(*result)
	assert.Equal(t, rArray[0].Status, "CREATE_IN_PROGRESS")
	assert.Equal(t, rArray[0].LogicalID, "MyLambdaRole")
	assert.Equal(t, rArray[1].Status, "CREATE_COMPLETE")
}
