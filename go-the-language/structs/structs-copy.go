package main

import (
	"fmt"
)

type Supe struct {
	name     string
	strength int
}

func main() {
	var shazam Supe
	shazam.name = "shazam"
	shazam.strength = 300
	//begin
	son := shazam
	son.name = "shazamson"
	fmt.Println(son)
	son.strength = 200
	fmt.Println(son)
	fmt.Println(shazam)
	//end
}
