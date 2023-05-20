package main

import (
	"fmt"
	"os"
	"table"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const tableName = "barjokes"

func main() {

	fmt.Printf("Creating table %v \n", tableName)
	err := table.CreateTable(aws.String(tableName))
	if err != nil {
		fmt.Printf("Error creating table %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Waiting for table to exist")
	err = table.Wait(aws.String(tableName))
	if err != nil {
		fmt.Printf("Error waiting for table to exist %v \n", err)
		os.Exit(2)
	}
	fmt.Printf("Table created \n")
	fmt.Printf("Fill table \n")
	err = table.FillTable(aws.String(tableName))
	if err != nil {
		fmt.Printf("Error filling table %v \n", err)
		os.Exit(3)
	}
	fmt.Printf("Table filled \n")

}
