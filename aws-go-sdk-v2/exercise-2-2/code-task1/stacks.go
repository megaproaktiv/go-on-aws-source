package crlist

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"golang.org/x/exp/slog"
)

// begin client
var Client *cloudformation.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	Client = cloudformation.NewFromConfig(cfg)

}

//end client

// begin type
// Type for holding logicalid and status
type ResourceStatus struct {
	LogicalID string
	Status    string
}

//end type

// GetStatus get States of all Cfn Stacks
// begin func
func GetStatus(client *cloudformation.Client, stackname *string) (*[]ResourceStatus, error) {
	//end func
	//begin call
	states := &[]ResourceStatus{}
	// Get resource status for stack stackname
	parms := &cloudformation.DescribeStackResourcesInput{
		StackName: stackname,
	}
	resp, err := client.DescribeStackResources(context.Background(), parms)
	if err != nil {
		slog.Error("Error in getting stack status", err)
		return nil, err
	}
	//end call
	//begin resultloop
	for _, resource := range resp.StackResources {
		*states = append(*states, ResourceStatus{
			LogicalID: *resource.LogicalResourceId,
			Status:    string(resource.ResourceStatus),
		})
	}
	return states, nil
	//end resultloop

}
