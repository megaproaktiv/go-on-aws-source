package main

import (
	"fmt"
	"reflect"
)

func main() {

	rune := 'B'
	aString := "B"

	// See content, unicode and type
	fmt.Printf("Rune: %v , %c; Unicode: %U; Type: %s \n", rune, rune, rune, reflect.TypeOf(rune))
	fmt.Printf("String: %v , %c; Unicode: %U; Type: %s \n", aString, aString, aString, reflect.TypeOf(aString))

}
