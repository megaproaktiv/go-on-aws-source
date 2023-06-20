package main

import "fmt"


func main() {
  //begin
  fmt.Printf("2 arguments: %v\n",sum(1, 2))
  fmt.Printf("3 arguments: %v\n",  sum(1, 2, 3))
  nums := []int{1, 2, 3, 4}
  // fmt.Printf("4 arguments: %v\n",  sum(nums))
  // cannot use nums (variable of type []int) as int value in argument to sum
  fmt.Printf("4 arguments: %v\n",  sum(nums...))

  //end
}

//begin
func sum(nums ...int) int{
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
//end
