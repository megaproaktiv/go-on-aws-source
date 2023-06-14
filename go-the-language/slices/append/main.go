package main

import (
	"fmt"
)

func main() {
  //begin
  manyKnifes := make([]string, 3, 5)
  fmt.Printf("Length: %v\n", len(manyKnifes))          // Length: 3
  manyKnifes = append(manyKnifes, "katana", "bowie")
  fmt.Printf("Length: %v\n", len(manyKnifes))          // Lenght: 5
  //end
}
