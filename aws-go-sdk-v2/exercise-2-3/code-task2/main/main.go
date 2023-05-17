package main

//begin main
import (
	"fmt"
	"os"
	"s3share"
)

func main() {
	bucket := "dateneimer"
	from := "testdata/text.txt"

	s3share.Upload(s3share.Client,&from, &bucket)
	url, err := s3share.Share(s3share.Client, &from, &bucket)

	if err != nil {
		fmt.Printf("Problem with sharing: %s",err)
		os.Exit(1)
	}
	fmt.Println(url)

}
//end main
