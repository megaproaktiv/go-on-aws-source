//begin
package main

import(
	"stackcounter"
	"fmt"

)

func main(){
	count := stackcounter.Count()
	fmt.Println("Counting CloudFormation Stacks: ",count)
}
//end
