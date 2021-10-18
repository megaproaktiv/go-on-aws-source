package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("s3event.json")
	if err != nil {
		fmt.Println("File read error:", err)
	}
	fmt.Println(string(content))
}
