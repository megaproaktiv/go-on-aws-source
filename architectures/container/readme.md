# Fargate Container architecture

app - Gin GO application to run in a fargate container
infra - ApplicationLoadBalancedFargateService

Details see https://www.go-on-aws.com/

## Install

### Create Application

```bash
task build-app
```

Export Account number for CDK

Get Account number:

```bash
aws sts get-caller-identity
```

Export as CDK_DEFAULT_ACCOUNT

```bash
export CDK_DEFAULT_ACCOUNT=555555555555
```

Replace the number with your account number.

### Create Infrastructure and run

Using CDK V2

#### VPC

```bash
task deploy-vpc
```

#### Table

```bash
task deploy-table
```


#### Container

```bash
task deploy-container
```

#### Run




## Cleanup

Reverse order

```bash
task destroy-container
task destroy-table
task destroy-vpc
```
