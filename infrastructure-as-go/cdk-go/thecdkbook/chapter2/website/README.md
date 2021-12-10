# Changes for GO

- Main function moves in own directory "main"
- Update cdk.json


When you change the name of the file containing "main", you have to update `cdk.json`:

```json
   "app": "go mod download && go run main/main.go",
```

Usually go modules are *not* camelcase, this is to map the cdkbook examples easier
