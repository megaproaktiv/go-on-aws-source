package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var client *sns.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = sns.NewFromConfig(cfg)
}

func main() {
	topic := aws.String("go-topic")

	parms := &sns.CreateTopicInput{
		Name: topic,
	}

	results, err := client.CreateTopic(context.TODO(), parms)
	if err != nil {
		panic("sns error, " + err.Error())
	}

	fmt.Println(*results.TopicArn)
}
