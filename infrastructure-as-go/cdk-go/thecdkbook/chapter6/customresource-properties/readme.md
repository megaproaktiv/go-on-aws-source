# 6.1.1. Implementing custom resources using AWS CDK

Parameter for custom resources


```go
awscdk.NewCustomResource(this, aws.String("MyResource"), &awscdk.CustomResourceProps{
		ServiceToken:         myFunction.FunctionArn(),
		Properties:           &map[string]interface{}{
			"Name" : "MyName",
			"Bucket" : myBucket.BucketName(),
		},
	})
``
