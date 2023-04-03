package main

import "fmt"

func main() {
	heroes := map[string]bool{
		"peter": true,
		"gwen":  false,
		"bruce": true}

	for i := range heroes {
		fmt.Println(i)
	}
	for i, supe := range heroes {
		fmt.Println(i, supe)
	}
}
