package main

import "fmt"

func main() {
	//begin
	heroes := map[string]bool{
		"peter": true,
		"gwen":  false,
		"bruce": true,
	}
	for i, v := range heroes {
		fmt.Printf("Name: %v, Bool: %v \n",i, v)
	}
	//end
}
