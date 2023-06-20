package main

import (
	"fmt"
)

//begin
type CustomerInfo struct{
  Vip bool
  Balance float64
}

//end

//begin
func main() {
	name := "Alice"
	ci := customerInfo(name)
	if ci.Vip {
	// ...
	//end
	 fmt.Printf("%v is a vip.\n", name)
	}else {
	 fmt.Printf("%v is no vip.\n",name)
	}
	fmt.Printf("The account balance is: %.2f.\n",ci.Balance)
//begin
}
//end

//begin function

func customerInfo(name string) CustomerInfo {
    returnedCI := CustomerInfo{
      Vip: true,
      Balance: 1000.00,
    }
    return returnedCI
}
//end function
