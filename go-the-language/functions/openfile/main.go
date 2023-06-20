package main

import (
  "os"
  "fmt"
  "log"
)



func main() {

    // Only open
    //begin defer
    f, err := os.Open("data.txt")
    if err != nil {
      log.Fatal(err)
    }
    defer f.Close()
    //end defer
    // Read data from f
    // ...

    //read all, do`t have to close
    //begin read
    dat, err := os.ReadFile("data.txt")
    if err != nil {
      log.Fatal(err)
    }
    fmt.Print(string(dat))
    //end read


}
