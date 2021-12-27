package main

import (
	"dbstack"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)


func main() {
	app := awscdk.NewApp(nil)

	dbstack.NewDbstackStack(app, "DevDb", &dbstack.DbstackStackProps{
		StackProps:   awscdk.StackProps{},
		VpcID:        aws.String("vpc-2f09a348"),
		InstanceType: aws.String("t3.micro"),
	})
	// RDS instance type would be "db.t3.micro", see https://aws.amazon.com/rds/instance-types/
	
	dbstack.NewDbstackStack(app, "DevDb", &dbstack.DbstackStackProps{
		StackProps:   awscdk.StackProps{},
		VpcID:        aws.String("vpc-abcd0123"),
		InstanceType: aws.String("r5.xlarge"),
	})
	// RDS instance type would be "db.r5.large", see https://aws.amazon.com/rds/instance-types/
	// EC2 instance types: https://aws.amazon.com/ec2/instance-types/

	app.Synth(nil)
}
