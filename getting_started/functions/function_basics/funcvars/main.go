package main

import ( "fmt")

func main(){
	var operations func(int,int) int
	operations = func(a int, b int)  int{
		return a+b
	}

	c := operations(1,2)
	fmt.Println(c)

	operations = func(a int, b int)  int{
		return a-b
	}

	c = operations(1,2)
	fmt.Println(c)


}

