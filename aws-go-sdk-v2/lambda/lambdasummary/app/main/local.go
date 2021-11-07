//+build local
package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"fmt"
	"lambdasummary"
)


func main() {
	accounts  := []*string{ aws.String("795048271754") }
	regions := []*string{
		aws.String("eu-central-1"), 
		aws.String("eu-west-1"),
		aws.String("us-east-1"),
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

	fmt.Println(content)

}


