package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	aStringPointer := aws.String("Hi")
	fmt.Println(aStringPointer)
	fmt.Println(*aStringPointer)
}
