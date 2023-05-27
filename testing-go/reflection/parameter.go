package awsmockdemo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var Client *ssm.Client


func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
			panic("configuration error, " + err.Error())
	}
	Client = ssm.NewFromConfig(cfg)
}

//begin testee
func GetTableName(client *ssm.Client) *string {
	
	parms := &ssm.GetParameterInput{
		Name: aws.String("/go-on-aws/table"),
	}
	resp, err := client.GetParameter(context.TODO(), parms)
	if err != nil {
		panic("ssm error, " + err.Error())
	}
	value := resp.Parameter.Value
	return value
}
//end testee