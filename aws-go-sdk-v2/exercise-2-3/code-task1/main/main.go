package main

import (
	"fmt"
	"os"
	"s3share"
)
//begin main
func main() {
	// replace "dateneiner" with your bucket name
	bucket := "dateneimer"
	from := "testdata/text.txt"
	err := s3share.Upload(s3share.Client,&from,&bucket)
	if err != nil {
		fmt.Printf("Problem with sharing: %s",err)
		os.Exit(1)
	}
}
//end main
