package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

var client *lambda.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
			panic("configuration error, " + err.Error())
	}
	client = lambda.NewFromConfig(cfg)

}

func main(){
	parms := &lambda.GetFunctionInput{
		FunctionName: aws.String("simple"),
	}

	resp, err := client.GetFunction(context.TODO(), parms)
	if err != nil{
		log.Fatal("Error getting function ", err)
	}
	e, err := json.MarshalIndent(resp, "", " ")
	if err != nil{
		log.Fatal("Error marshal to json function ", err)
	}
	fmt.Printf("Response: %v", string(e))
	
}
