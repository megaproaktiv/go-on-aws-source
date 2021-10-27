package main

import "fmt"

func main(){

	name := "Bob"
	say(name)
	fmt.Println("Name is still: ", name)

	sayRef(&name)
	fmt.Println("Name is now: ", name)
}

func say(name string){
	fmt.Println("Inside say: ",name);
	name = "Alice"
	fmt.Println("Inside say: ",name);
}

func sayRef(name *string){
	fmt.Println("Inside sayRef: ",*name);
	*name = "Alice"
	fmt.Println("Inside sayRef: ",*name);
}



