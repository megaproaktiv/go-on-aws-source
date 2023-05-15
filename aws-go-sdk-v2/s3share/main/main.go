package main

import (
	"flag"
	"s3share"
)

func main() {
	from := flag.String("file", "", "file to share")
	bucket := flag.String("bucket", "", "sharing bucket")

	flag.Parse()

	//call upload to s3
	s3share.Upload(*s3share.Client,from, bucket)


}
