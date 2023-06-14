package main

import "fmt"

func main() {
	//begin
	heroes := map[string]bool{
		"peter": true,
		"gwen":  false,
		"bruce": true}
	superman, ok := heroes["clark"]
	if ok {
		fmt.Print(superman, " is here!")
	} else {
		fmt.Print("Not here!")
	}
	//end
}
