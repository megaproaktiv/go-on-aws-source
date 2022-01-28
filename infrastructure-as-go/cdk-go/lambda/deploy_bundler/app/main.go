package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
        var n int64 = 31
        return fmt.Sprintf("Fibonacci(%d) = %d", n, fibo(n)), nil
}

func main() {
        lambda.Start(HandleRequest)
}

func fibo(n int64) int64 {
	if n <= 2 {
		return n - 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
