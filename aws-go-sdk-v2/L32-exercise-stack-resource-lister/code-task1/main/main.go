package main

import (
	"fmt"
	"os"

	"crlist"
)

func main() {
	//begin arguments
	argLength := len(os.Args[1:])
	if argLength == 0 {
		fmt.Printf("Please provide a stack name as an argument\n")
		os.Exit(1)
	} 
	stackName := os.Args[1]
	//end arguments
	//begin call
	resources, err := crlist.GetStatus(crlist.Client,&stackName)
	if err != nil {
		panic(err)
	}
	//end call

	//begin display
	fmt.Printf("%-32s %-32s \n", "Logical ID",  "Status")
	fmt.Printf("%-32s %-32s\n", "----------", "-----------")
	for _, resource := range *resources {
		fmt.Printf("%-32s %-32s\n", resource.LogicalID, resource.Status)
	}
	//end display

}