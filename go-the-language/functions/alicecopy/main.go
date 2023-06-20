package main

import (
	"fmt"
)
//begin
func main() {
	name := "Alice"
	say(name)
	say(name)
}

func say(name string) {
	fmt.Println(name)
	name = "white rabbit"
}
//end
