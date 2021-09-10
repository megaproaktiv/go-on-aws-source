package main

import(
	"stackcount"
	"fmt"

)

func main(){
	count := stackcount.Count()
	fmt.Println("Counting CloudFormation Stacks: ",count)
}