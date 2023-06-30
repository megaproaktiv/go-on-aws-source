//begin
// file: s3share.go
package s3share

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var Client *s3.Client

func init() {
	// create a s3 client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = s3.NewFromConfig(cfg)
}
//end