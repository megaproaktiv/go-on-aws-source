package main

import (
  "awsmockdemo"
  "fmt"
)

func main(){

  value := awsmockdemo.GetTableName(awsmockdemo.Client);
  fmt.Println("Value is: ", *value);

}