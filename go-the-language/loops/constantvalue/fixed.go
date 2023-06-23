package main

import "fmt"

func main() {
	//begin
	f := make([]func(), 3)
	for i := 0; i < 3; i++ {
		thisI := i // new variable each loop
		f[i] = func() {
			fmt.Println(thisI)
		}
	}
	fmt.Println("VariableLoop")
	for _, f := range f {
		f() // Calling different thisI
	}
	//end
}
