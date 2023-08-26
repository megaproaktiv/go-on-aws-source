package main

import "fmt"

func main() {
	// begin
	var x, y int
	x = 10
	y = 42
	m = min(x, y)
	fmt.Printf("Min %v\n", m)
	m = max(x, y, 10)
	fmt.Printf("Max %v\n", m)
	// end
}
