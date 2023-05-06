package main

import "fmt"

func main() {
        var varInt = 42
        var varString = "mega"
        var varFloat = 3.14
		varBool := true

        fmt.Printf("Value Int : %v\n", varInt)
        fmt.Printf("Type Int  : %T\n", varInt)

        fmt.Printf("Value string : %v\n", varString)
        fmt.Printf("Type string  : %T\n", varString)

        fmt.Printf("Value float : %v\n", varFloat)
        fmt.Printf("Type float  : %T\n", varFloat)

        fmt.Printf("Value bool : %v\n", varBool)
        fmt.Printf("Type bool  : %T\n", varBool)
}