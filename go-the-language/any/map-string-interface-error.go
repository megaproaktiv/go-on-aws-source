package main

import (
	"fmt"
)

func main() {

	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs":  any{"chicken", 1.75},
		"steak": true,
	}
	fmt.Printf("Food value: %v\n", foods)
}
