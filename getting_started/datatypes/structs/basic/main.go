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
	fmt.Println(shazam)
	shazam.name = "shazam"
	shazam.strength = 300
	fmt.Println(shazam)
	son := shazam
	fmt.Println(son)
	son.strength = 200
	fmt.Println(son)
	fmt.Println(shazam)
	daughter := new(Supe)
	daughter.strength = 400
	fmt.Println(daughter)

}
