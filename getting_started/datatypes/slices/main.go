package main

import "fmt"

func main() {
  manyKnifes := make([]string, 3,5)

  fmt.Println("How many:", len(manyKnifes))

  for _, slicer := range manyKnifes{
    fmt.Println((slicer))
  }
  
  manyKnifes = append(manyKnifes, "katana", "bowie")
  for _, slicer := range manyKnifes{
    fmt.Println((slicer))
  }
}
