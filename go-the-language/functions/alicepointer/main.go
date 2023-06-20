package main

import (
	"fmt"
)
//begin
func main() {
	name := "Alice"
	fmt.Println("1st")
	say(&name)
	fmt.Println("2nd")
	say(&name)
}

func say(name *string) {
	fmt.Println(*name)
	*name = "white rabbit"
}
//end
