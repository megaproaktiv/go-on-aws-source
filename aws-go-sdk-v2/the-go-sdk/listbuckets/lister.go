package listbuckets

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//begin init
var client *s3.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = s3.NewFromConfig(cfg)
}
//end init


func ListBuckets() ([]*string, error){
	bucketarray := make( []*string, 0,10)

	//begin sdkapicall
	res, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	//end sdkapicall

	//begin responseloop
	for _, bucket := range res.Buckets {
		bucketarray = append(bucketarray, bucket.Name)
	}
	//end responseloop

	return bucketarray, nil

}
