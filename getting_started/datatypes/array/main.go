package main

import "fmt"

func main() {
	var names [5]string
	names[0] = "hugo"
	names[2] = "melc"

	fmt.Println("My name is: " + names[0])
	fmt.Println("i dont know you,  " + names[1])

	// Range
	for i := 0; i < len(names); i++ {
		fmt.Println("Names ", i, ": ", names[i])
	}
	
	//Sub-parts
	fmt.Println("All names:", names)
	fmt.Println("Some names:", names[1:])
	fmt.Println("Some names:", names[:2])
	
	var namePointer [3]*string;
	namePointer[0] = &names[0]
	namePointer[1] = &names[2]
	
	for i, aPointer := range namePointer {
		fmt.Println( i, " : ", aPointer)
	}

	for i, aPointer := range namePointer {
		if aPointer != nil{
			fmt.Println( i, " : ", *aPointer)
		}
	}

}
