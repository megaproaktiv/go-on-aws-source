package main

import "fmt"

func main() {
	// begin
	x := 127
	c := max(1, 2.0, 10) // c == 10.0 (floating-point kind)
	fmt.Printf("Max %v\n", c)
	f := max(0, float32(x)) // type of f is float32
	fmt.Printf("Max %v\n", f)
	// end
}
