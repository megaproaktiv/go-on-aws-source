package main

import (
	"fmt"
)

type Supe struct {
	name     string
	strength int
}

func main() {
	//begin
	var shazam Supe
	shazam.name = "shazam"
	shazam.strength = 300
	fmt.Print(shazam)
	//end
}
