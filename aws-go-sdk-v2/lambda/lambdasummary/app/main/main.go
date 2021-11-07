//build cloud
package main

import (
	"context"
	"fmt"
	"lambdasummary"
	"strings"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/lambda"

	paddle "github.com/PaddleHQ/go-aws-ssm"
)

type MyEvent struct {
	Name string `json:"name"`
}

var configuration *paddle.Parameters

func init(){
	pmstore, err := paddle.NewParameterStore()
	if err != nil {
		log.Fatal("Cant connect to Parameter Store")
	}
	//Requesting the base path
	configuration, err = pmstore.GetAllParametersByPath("/showfunctions/", true)
	if err!=nil{
		log.Fatal("Can not get Parameter Store")
	}
	
}


func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	
	accounts := strings.Split(configuration.GetValueByName("accounts"), ",")
	
	regions := []string{
		"eu-central-1",
		"eu-west-1",
		"us-east-1",
	}

	
	summaries , err := lambdasummary.Collect(accounts, regions)

	infoPageData := &lambdasummary.FunctionListPageData{
		PageTitle: "Userlist",
		Summaries: summaries,
	}

	if err != nil {
		fmt.Println("Error: ",err)
	}

	content := lambdasummary.Render(*infoPageData)

	
	return content, nil
}

func main() {
	lambda.Start(HandleRequest)
}


