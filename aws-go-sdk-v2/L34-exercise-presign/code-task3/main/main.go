package main

import (
	"flag"
	"fmt"
	"os"
	"s3share"
)

func main() {
	//begin cli
	from := flag.String("file", "", "file to share")
	bucket := flag.String("bucket", "", "sharing bucket")
	flag.Parse()
	//end cli

	//call upload to s3
	//begin main
	s3share.Upload(s3share.Client,from, bucket)
	url, err := s3share.Share(s3share.Client, from, bucket)
	//end main
	// handle error
	if err != nil {
		fmt.Printf("Problem with sharing: %s",err)
		os.Exit(1)
	}

	fmt.Println(url)

}
