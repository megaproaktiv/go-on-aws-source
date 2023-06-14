package main

import "fmt"

func main() {
    //begin
    names := [...]string{"me","you", "doubleyou"}
    l := len(names)
    for i := 0; i < l; i++ {
            fmt.Println("Names ", i, ": ", names[i])
    }
    //end
}
