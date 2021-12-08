package main

import (
  "reflection"
  "fmt"
)

func main(){

  value := reflection.GetTableName(reflection.Client);
  fmt.Println("Value is: ", *value);

}