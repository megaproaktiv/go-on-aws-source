package main

import "fmt"

func main() {
	//begin
	for i := 3; i < 10 ; i++ {
		if i == 7  {
		  fmt.Println("You are lucky")
		  continue
		}
		if i % 2 == 0 {
		  fmt.Println(i," is even")
			continue
		}
		if i % 2 == 1 {
		  fmt.Println(i ," is odd")
		}
	}
	//end
}
