package main

import (
	"fmt"
	"os"

	"crlist"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength == 0 {
		fmt.Printf("Please provide a stack name as an argument\n")
	} 
	stackName := os.Args[1]
	resources, err := crlist.GetStatus(crlist.Client,&stackName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%-32s %-32s \n", "Logical ID",  "Status")
	fmt.Printf("%-32s %-32s\n", "----------", "-----------")
	
	for _, resource := range *resources {
		fmt.Printf("%-32s %-32s\n", resource.LogicalID, resource.Status)
	}

}