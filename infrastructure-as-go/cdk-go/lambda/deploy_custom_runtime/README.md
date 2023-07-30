# GO Lambda custom runtime

infra - Lambda with custom runtime al2
app - simple hello world go

## Install

### Configure

`infra/Taskfile.yml`

- Set CDK version for version pinning
- set AWS profile you use

```yaml
vars:
  version: 2.89.0
  PROFILE: letstrain
```

### Deploy

 * `task app:build` build lambda function
 * `task infra:deploy`         deploy stack without asking

 * `task -l `      see all possible tasks
 

## Sample walk-trough

1) Build

```bash
task app:build
  adding: bootstrap (deflated 59%)
```

2) Deploy

```bash
 task infra:deploy
Profile letstrain
task: [infra:deploy] npx  aws-cdk@2.89.0  deploy  --require-approval never --profile letstrain
✨  Synthesis time: 6.84s

hello-runtimeal2:  start: Building c716bac546c8b1e56143c1f8097e9306c6168bffa976ff3cf2f36e58eae0cc08:current_account-current_region
...
hello-runtimeal2: creating CloudFormation changeset...

 ✅  hello-runtimeal2

✨  Deployment time: 57.55s

Stack ARN:
arn:aws:cloudformation:eu-central-1:139008737997:stack/hello-runtimeal2/25be6710-2ea9-11ee-9dd9-0aa7c33ef562

✨  Total time: 64.39s
```

3) Destroy

```bash
task infra:destroy
Profileletstrain
task: [infra:destroy] npx  aws-cdk@2.89.0  destroy --force --profile letstrain
hello-runtimeal2: destroying... [1/1]

 ✅  hello-runtimeal2: destroyed
 ```
