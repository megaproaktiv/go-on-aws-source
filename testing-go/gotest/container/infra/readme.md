# Fargate Infrastructure


You need a VPC.

Check IDs with

```bash
aws ec2 describe-vpcs --query "Vpcs[].VpcId" --output table
```

Store this id in Systems Manager Parameter `/network/basevpc`

Check if its there with:

```bash
aws ssm get-parameter --name "/network/basevpc"
```

Output if its ok:


```json
{
    "Parameter": {
        "Name": "/network/basevpc",
        "Type": "String",
        "Value": "vpc-0cd99476a9a9c8579",
        "Version": 1,
        "LastModifiedDate": "2021-11-24T08:12:14.625000+01:00",
        "ARN": "arn:aws:ssm:eu-central-1:555555555555:parameter/network/basevpc",
        "DataType": "text"
    }
}
```

If not put the id into the store with:

```bash
aws ssm put-parameter --name "/network/basevpc" --type "String" --value "vpc-0cd99476a9a9c8579"
```

Replace the value  "vpc-0cd99476a9a9c8579" with your vpc ID.
