package main

import "fmt"

func main() {
	//begin
	heroes := map[string]bool{
		"peter": true,
		"gwen":  false,
		"bruce": true}
	for i := range heroes {
		fmt.Println(i)
	}
	//end
}
