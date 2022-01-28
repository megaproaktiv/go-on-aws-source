import { aws_dynamodb as dynamodb, Stack, StackProps, Tags } from 'aws-cdk-lib';
import { Construct } from 'constructs';

export class DynamodbcdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    const table = new dynamodb.Table( this, 'table',{
      partitionKey: { name: 'Username', type: dynamodb.AttributeType.STRING },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      tableName: "UserTableCDK",
    })

    Tags.of(table).add('Name','Username')
  }
}
