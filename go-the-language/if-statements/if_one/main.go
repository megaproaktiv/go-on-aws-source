package main

import (
        "fmt"
)

func main() {
//begin
        comment := map[string]string{
                "route53":       "Love it",
                "s3":            "",
                "ec2":           "oh!",
                "Snowball Edge": "Urban Dictionary both of those words.",
                "fargate":       "Sounds clever, but nobody will know what it does.",
                "graviton":      "Terrific if the next word in the product name is 'bomb'",
        }

        // Uncomment one of these three lines
        name := "ec2"
        // name := "route53"
        // name := "s3"

        if len(comment[name]) > 4 {
                fmt.Println("Clever comment for: ", name, " - ", comment[name])
        } else if len(comment[name]) > 0 {
                fmt.Println("Short comment for: ", name)
        } else {
                fmt.Println("No comment for: ", name)
        }
//end
}
