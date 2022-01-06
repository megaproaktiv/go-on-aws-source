# Input

Put VPCId of non default vpc in parameter store:


```bash
aws ssm put-parameter --name "/cdkbook/vpcid" --type String --value "vpc-0918e36a818968b79"
```

Replace vpcid with your vpcid

Export your account id and the region:

```bash
export CDK_DEFAULT_ACCOUNT=12345678901
export CDK_DEFAULT_REGION=eu-central-1
```
