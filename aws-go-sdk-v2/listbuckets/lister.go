package listbuckets

//begin import
import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)
//end import

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

	//begin params
	parms := &s3.ListBucketsInput{}
	//end params
	//begin sdkapicall
	res, err := client.ListBuckets(context.TODO(), parms)
	//end sdkapicall

	//begin response
	if err != nil {
		return nil, err
	}
	for _, bucket := range res.Buckets {
		bucketarray = append(bucketarray, bucket.Name)
	}
	//end response

	return bucketarray, nil

}
