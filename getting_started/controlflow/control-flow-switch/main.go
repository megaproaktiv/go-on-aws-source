package main

import (
	"fmt"
	"strings"
)

func main() {
	comment := map[string]string{
		"route53":       "LOVE IT",
		"s3":            "",
		"ec2":           "oh!",
		"Snowball Edge": "Urban Dictionary both of those words.",
		"fargate":       "Sounds clever, but nobody will know what it does.",
		"graviton":      "Terrific if the next word in the product name is 'bomb'",
	}

	// Uncomment one on these three lines
	name := "route53"
	// name := "ec2"
	// name := "s3"
	l := len(comment[name])
	switch {
	case l > 4:
		fmt.Println("Clever comment for: ", name, " - ", comment[name])
	case l > 0:
		fmt.Println("Short comment for: ", name)
	default:
		fmt.Println("No comment for: ", name)
	}
	
	switch l {
	case 7:
		fmt.Println("Comment with 7 characters: ", name, " - ", comment[name])
	case 0:
		fmt.Println("No comment for: ", name)
	}

	switch u := allUppercase(comment[name]); u {
	case true:
		fmt.Println("You dont have to shout")
	case false: 
		fmt.Println("ok")
	default:
		fmt.Println("This will never happen")
	}


}

func allUppercase(a string) bool{
	return strings.ToUpper(a) == a
}