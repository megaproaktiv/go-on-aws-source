package stackcounter

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var client *cloudformation.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        panic("unable to load SDK config, " + err.Error())
	}
	client = cloudformation.NewFromConfig(cfg);
}

//begin logic
func Count() (int){
	input := &cloudformation.DescribeStacksInput{}
	resp, _ := client.DescribeStacks(context.TODO(), input)
	count := len(resp.Stacks)
	return count
}
//end logic
