package showtable

import (
	// "context"

	// "github.com/aws/aws-sdk-go-v2/config"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)


func GetTableName(client *ssm.Client) *string {
	parms := &ssm.GetParameterInput{
		Name: aws.String("devopenspacetable"),
	}
	resp, err := client.GetParameter(context.TODO(), parms)
	if err != nil {
		panic("ssm error, " + err.Error())
	}
	value := resp.Parameter.Value
	return value
}