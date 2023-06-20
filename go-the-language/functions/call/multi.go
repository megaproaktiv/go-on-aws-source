package main

import (
	"fmt"
)
func main() {
  //begin
	worthy := creditworthy("Baltic Avenue", 1000 )
	fmt.Println("Customer is creditworthy:",worthy)
	worthy = creditworthy("Park Place", 10000 )
	fmt.Println("Customer is creditworthy:",worthy)

	//end
}

//begin
func creditworthy(street string, wealth float64 ) bool {
    if street == "New York Avenue" {
      return true
    }
    if wealth > 9000 {
      return true
    }
    return false
}
//end
