package main

import (
	"fmt"
	"listbuckets"
)

func main(){
	buckets,err := listbuckets.ListBuckets()

	if err != nil {
		fmt.Println("Error with listbuckets")
		panic(err)
		
	}
	for _, name := range buckets {
		fmt.Println(*name)
	}
}