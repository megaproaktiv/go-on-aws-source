package main

import (
	"fmt"
	"reflect"
)

func main() {
	//begin
	moreKnifes := []string{"Santoku", "Untility", "Paring"} // slice
	test(moreKnifes, "moreKnifes")

	manyKnifes := make([]string, 3, 5) //slice
	test(manyKnifes, "manyKnifes")

	allKnifes := [...]string{"Cleaver", "Steak", "Bread", "Tomato"} // Array
	test(allKnifes, "allKnifes")
	//end
}

func test(m interface{}, name string) {

	rt := reflect.TypeOf(m)
	switch rt.Kind() {
	case reflect.Slice:
		fmt.Printf("%v is a slice.\n", name)
	case reflect.Array:
		fmt.Printf("%v is an array.\n", name)
	default:
		fmt.Println("Is something else entirely.")
	}

}
