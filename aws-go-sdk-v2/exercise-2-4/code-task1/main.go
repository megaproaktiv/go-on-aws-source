package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/panjf2000/ants/v2"
)

var client *s3.Client
var uploader *manager.Uploader

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = s3.NewFromConfig(cfg)

}

func uploadOneObject(i interface{}) {
	n := i.(int32)
	key := fmt.Sprintf("test-go-pool-%d.md", n)

	path := "readme.md"
	file, err := os.Open(path)
	if err != nil {
		log.Println("Failed opening file", path, err)
	}
	defer file.Close()

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("dateneimer"),
		Key:    &key,
		Body:   file,
	})

	if err != nil {
		// Process error generically
		log.Println("Error:", err.Error())
		return
	}
	log.Printf("Loop %d\n", i)
}

func main() {

	const max = 30000
	const parallel = 50

	// Use the common pool.
	var wg sync.WaitGroup

	uploader = manager.NewUploader(client)

	p, _ := ants.NewPoolWithFunc(parallel, func(i interface{}) {
		uploadOneObject(i)
		wg.Done()
	})

	for i := 1; i < max; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
}
