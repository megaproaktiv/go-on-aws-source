package main

import "fmt"

func main() {
	// begin
	arrayInt := []int{1, 2, 3, 4, 6, 7, 8, 9}

	fmt.Printf("Length of arrayInt: %d\n", len(arrayInt))
	fmt.Printf("ArrayInt: %d\n", arrayInt)

	clear(arrayInt)
	fmt.Printf("Length of arrayInt: %d\n", len(arrayInt))
	fmt.Printf("ArrayInt: %d\n", arrayInt)
	// end
}
