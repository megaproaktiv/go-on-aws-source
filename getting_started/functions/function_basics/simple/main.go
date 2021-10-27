package main

import ( "fmt")

func main(){

	c := add(1,2)
	fmt.Println(c)
}

func add(a int, b int)  int{
	return a+b
}
