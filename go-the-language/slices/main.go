package main

import "fmt"

func main() {
	manyKnifes := make([]string, 3, 5)
	fmt.Printf("Length: %v\n", len(manyKnifes)) // Length: 3
	manyKnifes = append(manyKnifes, "katana", "bowie")
	fmt.Printf("Length: %v\n", len(manyKnifes)) // Lenght: 5

	allKnifes := []string{"Cleaver", "Steak", "Bread", "Tomato"} // Array
	fmt.Printf("All knifes  %v\n", allKnifes)
	someKnifes := allKnifes[1:2]                //slice
	fmt.Printf("Some knifes  %v\n", someKnifes) // [Steak]
	leftKnifes := allKnifes[:2]                 //slice
	fmt.Printf("Left knifes  %v\n", leftKnifes) //  [Cleaver Steak]
	leftKnifes = allKnifes[0:2]                 //slice
	fmt.Printf("Left knifes  %v\n", leftKnifes) //  [Cleaver Steak]
	rightKnifes := allKnifes[3:]
	fmt.Printf("Right knifes %v\n", rightKnifes) // [Tomato]
	rightKnifes = allKnifes[3:len(allKnifes)]
	fmt.Printf("Right knifes %v\n", rightKnifes) // [Tomato]
}
