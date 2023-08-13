package main

import "fmt"

func main() {
	// begin
	arrayInt := []int{1, 2, 3, 4, 6, 7, 8, 9}
	c := max(arrayInt) // c == 10.0 (floating-point kind)
	fmt.Printf("Max %v\n", c)
	// end
	// # command-line-arguments
	// ./minmax_noslice.go:8:11: invalid argument: arrayInt (variable of type []int) cannot be ordered
}
