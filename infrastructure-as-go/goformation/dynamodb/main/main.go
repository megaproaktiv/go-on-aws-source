package main

import (
	"fmt"
	"phantom"
)

func main(){
	content, err := phantom.Synth()
	if err != nil{
		panic(err)
	}
	fmt.Println(*content)
}