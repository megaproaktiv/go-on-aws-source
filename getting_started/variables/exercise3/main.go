package main

import "fmt"

var i = 42

func main() {
	i := 40
	for true {
		i := 43
		fmt.Println("Scope for:", i)
		i = 44
		fmt.Println("Scope for:", i)
		break
	}
	fmt.Println("Scope main:", i)
	outer()
}

func outer() {
	fmt.Println("Function: ", i)
}
