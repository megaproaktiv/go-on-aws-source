package showtable

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var ClientDDB *dynamodb.Client
var ClientSSM *ssm.Client


func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(lo *config.LoadOptions) error {
		lo.Region = "eu-central-1"
		return nil
	})
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	ClientDDB = dynamodb.NewFromConfig( cfg)
	ClientSSM = ssm.NewFromConfig( cfg)
}