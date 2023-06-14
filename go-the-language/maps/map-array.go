package main

import "fmt"

func main() {
	names := make([]string, 4)
	//begin
	names[1] = "hugo"
	fmt.Printf("First name is: %v\n", names[1])
	//end
}
