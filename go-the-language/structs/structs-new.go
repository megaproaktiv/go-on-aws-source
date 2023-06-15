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
	daughter := new(Supe)
	daughter.name = "not Hespera"
	daughter.strength = 400
	fmt.Println("Point: ", daughter)
	fmt.Println("Itself: ", *daughter)
	//end
}
