package main

import (
	"fmt"
)

func main() {
	hello := "hello"
	msg := &hello
	fmt.Println("Say ")
	unchangedSay(msg)
	fmt.Println("Unchanged: ", *msg)
	changedSay(msg)
	fmt.Println("Changed: ", *msg)
}

func unchangedSay(parm *string) {
	goodbye := "Wave goodbye"
	parm = &goodbye
}
func changedSay(parm *string) {
	goodbye := "Wave goodbye"
	*parm = goodbye
}
