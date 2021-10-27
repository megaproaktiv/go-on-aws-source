package main

import "fmt"

func main() {
	defer fmt.Println("exit")
	fmt.Println("start")
}
