# 2.2.3. Level 2 Constructs


```go
myHandler := awslambda.NewFunction(stack, aws.String("myHandler"), &awslambda.FunctionProps{})

myTable := awsdynamodb.NewTable(stack, aws.String("items"), &awsdynamodb.TableProps{})

myTable.GrantReadWriteData(myHandler);

```
