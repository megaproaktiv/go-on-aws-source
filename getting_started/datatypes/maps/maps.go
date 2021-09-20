package main

import "fmt"

func main() {
	heroes := []string{"peter", "gwen", "bruce"}

	for i := range heroes {
		fmt.Println(i)
	}
	for i, supe := range heroes {
		fmt.Println(i, supe)
	}
}
