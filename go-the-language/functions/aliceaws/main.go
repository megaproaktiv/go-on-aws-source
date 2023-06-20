package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"

)
//begin
func main() {
	name := "Alice"
	fmt.Println("1st")
	say(&name)
	fmt.Println("2nd")
	say(&name)
}

func say(name *string) {
	fmt.Println(*name)
	name = aws.String("white rabbit")
}
//end
