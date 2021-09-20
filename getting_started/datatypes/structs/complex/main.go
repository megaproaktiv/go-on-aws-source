package main

import (
	"fmt"
)

type Strength int

const (
	Weak Strength = iota
	Middle
	Strong
	Galactus
)

type Supe struct {
	Name       string
	Powerlevel Strength
}

func main() {
	var shazam Supe
	shazam.Name = "shazam"
	shazam.Powerlevel = Strong
	fmt.Println(shazam)

}
