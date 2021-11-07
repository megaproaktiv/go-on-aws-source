//build cloud local
package lambdasummary

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaInfo struct {
	Account string
	Region  string
	Name    string
	Runtime string
	Memory  int32
}

const max = 100

func ListFunctions(client *lambda.Client, account string, region string) []*LambdaInfo {

	lambdaFunctions := []*LambdaInfo{}

	params := &lambda.ListFunctionsInput{
		MaxItems: aws.Int32(int32((max))),
	}

	result, err := client.ListFunctions(context.TODO(), params, func(options *lambda.Options) {
		options.Region = region
	})
	if err != nil {
		fmt.Println("Got an error retrieving functions:")
		fmt.Println(err)
		return nil
	}

	for _, item := range result.Functions {
		info := &LambdaInfo{
			Name:    *item.FunctionName,
			Runtime: string(item.Runtime),
			Memory:  *item.MemorySize,
			Account: account,
			Region:  region,
		}
		lambdaFunctions = append(lambdaFunctions, info)
	}
	return lambdaFunctions
}
