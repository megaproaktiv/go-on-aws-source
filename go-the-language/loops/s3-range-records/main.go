package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	buckets, err := ListBuckets()

	if err != nil {
		fmt.Println("Error with listbuckets")
		panic(err)

	}
	for _, name := range buckets {
		fmt.Println(*name)
	}
}

var client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = s3.NewFromConfig(cfg)
}

func ListBuckets() ([]*string, error) {
	buckets := make([]*string, 0)

	res, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	//begin
	for _, bucket := range res.Buckets {
		buckets = append(buckets, bucket.Name)
	}
	fmt.Print("Buckets: ", buckets)
	//end
	return buckets, nil

}
