package main

import (
  "fmt"
  "strings"
)
//begin
func bothCases(a string) (lower, upper string) {
  lower = strings.ToLower(a)
  upper = strings.ToUpper(a)
  return
}

func main() {
	fmt.Println(bothCases("Alice"))
}
//end
