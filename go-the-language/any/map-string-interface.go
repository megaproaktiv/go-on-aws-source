package main

import (
	"fmt"
)

func main() {

	//begin
	foods := map[string]interface{}{
		"ham": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	fmt.Printf("Food value: %v", foods)
	//end
}
