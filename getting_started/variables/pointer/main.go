package main

import "fmt"

func main() {
	coat := "I am a coat"
	receipt := &coat

	receipt2 := &coat

	fmt.Println(*receipt)
	fmt.Println(*receipt2)
	
	var anotherCoat string
	anotherCoat = "I am the red coat"
	
	var redReceipt *string
	redReceipt = &anotherCoat
	fmt.Println(*redReceipt)
	
	redReceipt = &coat
	fmt.Println(*redReceipt)

}
