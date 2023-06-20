package main

import (
	"fmt"
)
func main() {
  //begin
	name := "Alice"
	vip, balance := customerInfo(name)

	//end
	if vip {
	 fmt.Printf("%v is a vip.\n", name)
	}else {
	 fmt.Printf("%v is no vip.\n",name)
	}
	fmt.Printf("The account balance is: %.2f.\n",balance)
}

//begin function
func customerInfo(name string) (bool, float64) {
    vip := true
    balance := 1000.00
    return vip, balance
}
//end function
