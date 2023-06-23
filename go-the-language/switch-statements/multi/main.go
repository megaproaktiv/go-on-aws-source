package main

import "fmt"

func main() {
	// assume that `tag` is a variable containing the tag value
	//begin
	number_below_10 := 2
	switch number_below_10 {
	default:
		fmt.Println("I don´t know")
	case 2, 3, 5, 7:
		fmt.Println("It´s a prime")
	case 0,1,4,6,8,9:
		fmt.Println("It´s not a prime")
	}
	//end
}
