# Deploy goformation stacks with cli app

This creates a DynamoDB Table with goformation/CloudFormation

## Commands

- help : usage
- deploy : deploy cloudformation
- status : show resources
- destroy : destroy cloudformation
 
## Example

```bash
go run main/main.go deploy
```

Or build first

```bash
go build -o dist/cfn main/main.go
```

Then

```bash
dist/cfn deploy
```