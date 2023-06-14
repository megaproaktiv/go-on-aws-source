package main

import "fmt"

func main(){
//begin parts
 names := [...]string{"huey","dewey", "louie"}
 fmt.Println("All names:", names)
 fmt.Println("Some names:", names[1:])
 fmt.Println("Some names:", names[:2])
//end
}
