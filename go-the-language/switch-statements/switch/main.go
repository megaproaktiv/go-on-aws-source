package main

import (
	"fmt"
	"strings"
)

func main() {
	//begin comments
	comment := map[string]string{
		"route53":       "Love it",
		"s3":            "",
		"ec2":           "oh!",
		"Snowball Edge": "Urban Dictionary both of those words.",
		"fargate":       "Sounds clever, but nobody will know what it does.",
		"graviton":      "Terrific if the next word in the product name is 'bomb'",
	}
	//end comments
	// Uncomment one on these three lines
	// name := "ec2"
	// name := "s3"
	//begin switch
	name := "route53"
	l := len(comment[name])
	switch {
	case l > 4:
		fmt.Println("Clever comment for: ", name, " - ", comment[name])
	case l > 0:
		fmt.Println("Short comment for: ", name)
	default:
		fmt.Println("No comment for: ", name)
	}
	//end switch

	//begin initialized
	switch u := allUppercased(comment[name]); u {
	case true:
		fmt.Println("You dont have to shout")
	case false:
		fmt.Println("ok")
	default:
		fmt.Println("This will never happen")
	}
	//end initialized


}

//begin function
// Check whether all character in a string are uppercase
func allUppercased(a string) bool {
	return strings.ToUpper(a) == a
}
//end function
