package main

import "fmt"

func main() {
	//begin
	i := 1
	for {
		fmt.Print(i, "-")
		i++
		if i == 10 {
			break
		}
	}
	//end
}
