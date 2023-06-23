package main

import "fmt"

func main() {
	//begin
	f := make([]func(), 3)
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	// now i is 3
	fmt.Println("VariableLoop")
	for _, f := range f {
		f() // calling i
	}
	//end
}
