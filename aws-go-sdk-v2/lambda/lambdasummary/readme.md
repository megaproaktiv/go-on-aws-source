# Build Organisation wide view

Iteration through accounts and regions here with "Lambda.ListFunctions".

You need a main account, in this example "11111111111" and some accounts to be listed the member accounts


## Directories

### infra

CDK V2 infrastructure

### app

GO SDK V2 app

## Installation

- Install task, GO, CDK2
See chapters in https://www.go-on-aws.com/ for details

### 1 Build Lambda function

In `app`:

```bash
task build
```

This creates a "main.zip" in `app/dist".

### 2 Configure member accounts

In main account, create Systems Manager Parameter named `/showfunctions/accounts` with all member accounts, comma seperated, to whitespaces.
This configures account fou the lambda and for the cdk

### 3 Deploy Lambda resource

In `app` in main account:

```bash
cdk deploy
```

Where cdk is an alias to `cdk='npx cdk@v2.0.0-rc.27'` or an globaly installed CDK2

### 4 Deploy CrossAccount role

Configure main account in `infra/policy/template.yaml`

Replace in line 10 the account number with your *main* account number

```yaml
              AWS: 'arn:aws:iam::111111111111:role/showfunctionsrole'
```

Deploy stack with

```bash
task trust
```

