package main

import "fmt"

func main() {

	// var declares a variable
	// and sets the type
	var varInt = 42
	var varString = "mega"
	var varFloat = 3.14

	// fmt "%v" just print almost everything
	// fmt "%T" prints the type
	fmt.Printf("Value Int : %v\n", varInt)
	fmt.Printf("Type Int  : %T\n", varInt)
	fmt.Printf("Type String  : %T\n", varString)
	fmt.Printf("Type Float  : %T\n", varFloat)

}
