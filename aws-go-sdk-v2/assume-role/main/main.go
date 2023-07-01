package main

import (
	"assume"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Main function
func main() {

	// Call the function
	user1, err := assume.Who(assume.Client)
	//handle errors
	if err != nil {
		log.Println("configuration error, " + err.Error())
		os.Exit(1)
	}
	account, err := assume.Where(assume.Client)
	if err != nil {
		log.Println("configuration error, " + err.Error())
		os.Exit(1)
	}
	role := "arn:aws:iam::" + *account + ":role/administrator"

	fmt.Printf("User: %s\n", *user1)
	var client2 sts.Client
	// Call the function
	cfg, err := assume.GetCfgSub(assume.Client, &role)
	//handle errors
	if err != nil {
		log.Println("configuration error, " + err.Error())
		os.Exit(2)
	}
	//begin client2
	client2 = *sts.NewFromConfig(*cfg)
	//end client2
	user2, err := assume.Who(&client2)
	//handle errors
	if err != nil {
		log.Println("configuration error, " + err.Error())
		os.Exit(3)
	}
	fmt.Printf("User: %s\n", *user2)
}
