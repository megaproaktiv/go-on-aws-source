package main

import "fmt"

func main() {
	lastnames := make(map[string]string)
	//begin
	lastnames["Hugo"] = "Boss"
	fmt.Printf("Last name is: %v\n", lastnames["Hugo"])
	//end
}
