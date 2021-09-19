package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 5
	var j int
	f := 12.1
	fmt.Printf("Int %v, Type: %s \n", i, reflect.TypeOf(i))
	fmt.Printf("Int %v, Type: %s \n", j, reflect.TypeOf(j))
	fmt.Printf("Float %v, Type: %s \n", f, reflect.TypeOf(f))

	// notOk := i + f
	var f2 float64 = float64(i)

	ok := f + f2
	fmt.Printf("Sum of f and f2: %v \n", ok)

}
