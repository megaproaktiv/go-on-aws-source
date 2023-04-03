package gfcli

import (
	"context"

	config "github.com/aws/aws-sdk-go-v2/config"
	cfnservice "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var Client *cfnservice.Client

// Client get a CloudFormation service client
func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	Client = cfnservice.NewFromConfig(cfg)
	InitLogger()		
	defer Logger.Sync()
}