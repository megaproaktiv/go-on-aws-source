package main

import "fmt"

func main() {
	//begin
	allKnifes := []string{"Cleaver", "Steak", "Bread", "Tomato"}
	for i, k := range allKnifes {
		fmt.Printf("Interator int: %v , value %v\n", i, k)
	}
	//end
}
