package main

import (
	"flag"
	"fmt"
	"os"
	"s3copy"
)

func main() {
	// Get file, bucket and refix as flags
	file := flag.String("file", "", "File to upload")
	bucket := flag.String("bucket", "", "S3 Bucket")
	prefix := flag.String("prefix", "", "Prefix")
	flag.Parse()

	if *file == "" || *bucket == "" || *prefix == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "Error: --file, --bucket and --prefix are required\n")
		os.Exit(1)
	}
	fmt.Printf("Uploading %s to %s/%s\n", *file, *bucket, *prefix)
	err := s3copy.Upload(s3copy.Client,  file, bucket, prefix)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}