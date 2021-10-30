# Deploy Lambda with Container Images

## Very quick walkthrough with [task](https://taskfile.dev)

Have installed

- [Docker](https://www.docker.com/)
- [task](https://taskfile.dev)
- [CDK V2](https://aws.amazon.com/cdk/)
- [jq](https://stedolan.github.io/jq/)
- [AWS CLI](https://aws.amazon.com/cli/)
- [GO](https://golang.org/)

See Steps

```bash
task -l
```

Do steps


```bash
task test
task deploy
task invoke
task show
task destroy
```



## Directory overview

This is for CDK V2

Directory | Content
---|---
app       | Simple Lambda as container build for arm
appx86    | Simple Lambda as container build for amd64 
infra     | CDK to deploy both

## Walktrough

1. Have local docker running

1. Change to infrastructure directory

    ```bash
    cd infra
    ```

1. Test

    ```bash
    go test
    ```

    Output:

    ```log
    PASS
    ok  	gograviton	2.599s
    ````

1. Deploy

    ```bash
    cdk deploy
    ```

    Output

    ```log
    lambda-go-arm: deploying...
    [0%] start: Publishing 89419431db330ace6f44d3dc25b3f689ca0bc06818811eee859abd8372035574:current_account-current_region
    [33%] success: Published 89419431db330ace6f44d3dc25b3f689ca0bc06818811eee859abd8372035574:current_account-current_region
    [33%] start: Publishing 5d242b86badb0a538e20c7692195f4e98046f1705c9f3c684b7c472a3f1aa470:current_account-current_region
    [66%] success: Published 5d242b86badb0a538e20c7692195f4e98046f1705c9f3c684b7c472a3f1aa470:current_account-current_region
    [66%] start: Publishing 88295ac237c1e453b47088f23f1c19ba5d206c22194d1befbd7eabfa9825cde6:current_account-current_region
    [100%] success: Published 88295ac237c1e453b47088f23f1c19ba5d206c22194d1befbd7eabfa9825cde6:current_account-current_region
    lambda-go-arm: creating CloudFormation changeset...
    [··························································] (0/6)

    07:46:50 | CREATE_IN_PROGRESS   | AWS::CloudFormation::Stack | lambda-go-arm
    ...
     ✅  lambda-go-arm

    Stack ARN:
    arn:aws:cloudformation:eu-central-1:555544443333:stack/lambda-go-arm/c3cc0590-3944-11ec-a61b-0ad372775e10
    ````

1. Call Lambda

    x86

    ```bash
    aws lambda invoke --function-name hellodockerx86  --payload fileb://testdata/event.json testdata/lambda.out
    ```

    arm

     

    ```bash
    aws lambda invoke --function-name hellodockerarm  --payload fileb://testdata/event.json testdata/lambda.out
    ```


    Response of the invoke api call is

    ```json
    {
    "StatusCode": 200,
    "ExecutedVersion": "$LATEST"
    }
    ```

    Response of lamba is:

    ```bash
    cat testdata/lambda.out
    ```

    Output

    ```log
    "Hiho doit!"
    ```

  

    Lambda logs: 

    ## arm

    Cold start

    ```log
    START RequestId: 6669704e-83f6-4358-b51f-d30be90207f1 Version: $LATEST
    END RequestId: 6669704e-83f6-4358-b51f-d30be90207f1
    REPORT RequestId: 6669704e-83f6-4358-b51f-d30be90207f1	Duration: 1.32 ms	Billed Duration: 548 ms	Memory Size: 1024 MB	Max Memory Used: 15 MB	Init Duration: 546.21 ms	
    ```

    Warm start

    ```log
    START RequestId: cd6ff89b-7298-4f38-8f9b-df2824a256da Version: $LATEST
    END RequestId: cd6ff89b-7298-4f38-8f9b-df2824a256da
    REPORT RequestId: cd6ff89b-7298-4f38-8f9b-df2824a256da	Duration: 0.77 ms	Billed Duration: 1 ms	Memory Size: 1024 MB	Max Memory Used: 15 MB	
    ```

    ## amd64

    Cold start

    ```bash
    START RequestId: 2cb952a1-9765-4d41-a21e-b03c265cc7dc Version: $LATEST
    END RequestId: 2cb952a1-9765-4d41-a21e-b03c265cc7dc
    REPORT RequestId: 2cb952a1-9765-4d41-a21e-b03c265cc7dc	Duration: 1.38 ms	Billed Duration: 796 ms	Memory Size: 1024 MB	Max Memory Used: 16 MB	Init Duration: 793.69 ms	
    ```

    ```log
    START RequestId: addabc9f-e1d3-42fc-b3c3-03f2ed40e72a Version: $LATEST
    END RequestId: addabc9f-e1d3-42fc-b3c3-03f2ed40e72a
    REPORT RequestId: addabc9f-e1d3-42fc-b3c3-03f2ed40e72a	Duration: 0.78 ms	Billed Duration: 1 ms	Memory Size: 1024 MB	Max Memory Used: 16 MB	
    ```

1. See created resources

    ## CloudFormation

    With [cdkstat](https://github.com/megaproaktiv/cdkstats)

    ```bash
    cdkstat lambda-go-arm
    ```

    Output

    ```log
    Logical ID                       Pysical ID                       Type                             Status
    ----------                       ----------                       -----------                      -----------
    CDKMetadata                      c3cc0590-3944-11ec-a61b-0ad3727  AWS::CDK::Metadata               CREATE_COMPLETE
    RegisterHandlerAmd9BBD5506       hellodockerx86                   AWS::Lambda::Function            CREATE_COMPLETE
    RegisterHandlerAmdServiceRole5F  lambda-go-arm-RegisterHandlerAm  AWS::IAM::Role                   CREATE_COMPLETE
    RegisterHandlerArm9EEB6A7A       hellodockerarm                   AWS::Lambda::Function            CREATE_COMPLETE
    RegisterHandlerArmServiceRole9D  lambda-go-arm-RegisterHandlerAr  AWS::IAM::Role                   CREATE_COMPLETE
    ```

    ## ECR - Container Registry

    ```bash
    aws ecr describe-repositories
    ```

    ```json
    {
    "repositories": [
        {
            "repositoryArn": "arn:aws:ecr:eu-central-1:5555:repository/cdk-hnb659fds-container-assets-5555-eu-central-1",
            "registryId": "5555",
            "repositoryName": "cdk-hnb659fds-container-assets-5555-eu-central-1",
            "repositoryUri": "5555.dkr.ecr.eu-central-1.amazonaws.com/cdk-hnb659fds-container-assets-5555-eu-central-1",
            "createdAt": "2021-08-02T10:44:53+02:00",
            "imageTagMutability": "MUTABLE",
            "imageScanningConfiguration": {
                "scanOnPush": false
            },
            "encryptionConfiguration": {
                "encryptionType": "AES256"
            }
        }
    ]
    }
    ```

    ### Images

    With the manualy extraced repositoryName do

    ```bash
    aws ecr list-images --repository-name cdk-hnb659fds-container-assets-5555-eu-central-1
    ```

    Or with [jq](https://stedolan.github.io/jq/)

    ```bash
    export name=`aws ecr describe-repositories | jq '.repositories[0].repositoryName' | tr -d '"' `
    aws ecr list-images --repository-name $name
    ```

1. Destroy

    ```bash
    cdk destroy
    ```