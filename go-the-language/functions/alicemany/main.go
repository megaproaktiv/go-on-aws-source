package main

import (
	"fmt"
)
//begin
func main() {
	name := "Alice"
	tell(name)
//end
	say(name)
}

func say(name string) {
	fmt.Println(name)
}
