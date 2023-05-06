// begin package
package instancelister_test

import (
	"context"
	"encoding/json"
	"fmt"
	"instancelister"
	"os"
	"testing"
	"gotest.tools/assert"
	
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/megaproaktiv/awsmock"
)

func TestListInstances(t *testing.T) {

	describeInstancesMock := func(ctx context.Context, params *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error){	
		testfile := "testdata/describe-instances.json"
		data, err := os.ReadFile(testfile)
		if err != nil {
			fmt.Println("File reading error: ", err)
		}
		out := &ec2.DescribeInstancesOutput{}
		err =  json.Unmarshal(data, out)
		if err != nil {
			t.Error(err)
		}
		return out, nil
	}

	mockCfg := awsmock.NewAwsMockHandler()
	mockCfg.AddHandler(describeInstancesMock)
	client := ec2.NewFromConfig(mockCfg.AwsConfig())
	instances, err := instancelister.ListInstances(client)
	assert.NilError(t, err)
	assert.Equal(t, len(instances), 1)

}
