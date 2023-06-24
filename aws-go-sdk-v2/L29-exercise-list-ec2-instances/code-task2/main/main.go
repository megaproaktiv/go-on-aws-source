//begin
package main

import (
	"instancelister"
	"fmt"

)

func main() {
	instances, err := instancelister.ListInstances(instancelister.Client)
	if err != nil {
		fmt.Print("Error calling instancelister: ", err)
	}
	for i,instance := range instances {
		fmt.Printf("Instance number: %v, ID: %v, Status: %v \n",
			i+1,instance.Name,instance.Status,
		)

	}
}
//end
