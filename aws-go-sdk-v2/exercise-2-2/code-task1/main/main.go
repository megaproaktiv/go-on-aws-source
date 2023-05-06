package main

import (
	"fmt"
	"os"

	"crlist"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength == 0 {
		overview()
	}else{
		stackName := os.Args[1]
		resources,err := crlist.ReadStackDetail(&stackName)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-32s %-32s %-32s %-32s \n", "Logical ID", "Pysical ID", "Type" , "Status")
		fmt.Printf("%-32s %-32s %-32s %-32s\n", "----------", "----------", "-----------","-----------")

		for _, resource := range *resources {
			
			logicalID := FixedLengthString(*resource.LogicalResourceId)
			physicalID := FixedLengthString(*resource.PhysicalResourceId)
			rType := FixedLengthString(*resource.ResourceType)
			statusString := string(resource.ResourceStatus)
			rStatus := FixedLengthString(statusString)
			fmt.Printf("%s %s %s %s\n", logicalID, physicalID, rType, rStatus)


		}

	}
}
func overview() {
	LOCALONLY := "LOCAL_ONLY"

	remoteStacks := crlist.GetStatus()
	localCDKStackNames := crlist.ReadStacks()
	remoteStackNames := make([]string, 5)

	fmt.Printf("%-32s %-32s %-32s \n", "Name", "Status", "Description")
	fmt.Printf("%-32s %-32s %-32s \n", "----", "------", "-----------")
	// Remote State
	for i := range remoteStacks.Stacks {
		stack := remoteStacks.Stacks[i]
		remoteStackNames = append(remoteStackNames, *stack.StackName)
		name := FixedLengthString(*stack.StackName)
		status := FixedLengthString(string(stack.StackStatus))
		description := "-"
		if stack.Description != nil{
			description = FixedLengthString(string(*stack.Description))
		}
		if contains(localCDKStackNames, *stack.StackName) {
			fmt.Printf("%s %s %s\n", name, status, description)
		}
	}
	// Local only
	status := FixedLengthString(LOCALONLY)
	for _, nameLocal := range *localCDKStackNames {
		name := FixedLengthString(*&nameLocal)
		if !contains(&remoteStackNames, nameLocal) {
			fmt.Printf("%s %s\n", name, status)
		}
	}

}

// FixedLengthString some formatting
func FixedLengthString(str string) string {
	if len(str) > 31{
		str = str[0:31]
	}
	return fmt.Sprintf("%-32s", str)
}

// does slice contain key
func contains(stacks *[]string, stack string) bool {
	for _, cdkStack := range *stacks {
		theSame := (cdkStack == stack)
		if theSame {
			return true
		}
	}
	return false
}
