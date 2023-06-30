# Stack resources

## Setup test environment

```bash
task setup
```

## Show resources

```bash
task describe
```

## Create

Write go program
The output should be a markdown table with LogicalResourceId and ResourceStatus:

```
LogicalResourceId | ResourceStatus
:---|:---
MyLambdaRole | CREATE_IN_PROGRESS
MySNSTopic | CREATE_COMPLETE
```


## Teardown

```bash
task teardown
```
