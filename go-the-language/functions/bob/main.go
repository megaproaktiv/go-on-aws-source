package main

import (
	"errors"
	"fmt"
)
//begin return
func main() {
//end return
//begin call

	name := "Alice"
	say(name)

//end call
//begin return
	age, err := howOldis(name)
	if err != nil {
		fmt.Printf("Don't know the age of %v\n", name)
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("%v is %v years old.\n", name, age)
	}
}

//end return

//begin call
func say(name string) {
	fmt.Println(name)
}
//end call
//begin return
func howOldis(name string) (int, error) {
	if name == "Bob" {
		return 42, nil
	}
	return 0, errors.New("name unknown")
}
//end return
