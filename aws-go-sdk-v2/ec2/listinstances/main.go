package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var client *ec2.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = ec2.NewFromConfig(cfg)

}

func main() {
	parms := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(10),
	}
	result, err := client.DescribeInstances(context.TODO(),parms)

	if err != nil {
		fmt.Println("Error calling ec2: ",err)
		return
	}
	count := len(result.Reservations)
	fmt.Println("Instances: ",count)

	for i, reservation := range result.Reservations {
		for k, instance := range reservation.Instances {
			fmt.Println("Instance number: ",i,"-",k	, "Id: ", instance.InstanceId)
		}
	}
}
