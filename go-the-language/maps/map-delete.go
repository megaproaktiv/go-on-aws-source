package main

import "fmt"

func main() {
	//begin
	weights := make(map[string]float64)
	weights["conan"] = 116.5
	fmt.Println("Conan: ", weights["conan"])
	fmt.Println("Arnold: ", weights["arnold"])
	delete(weights, "conan")
	fmt.Println("Conan has left: ", weights["conan"])
	//end
}
