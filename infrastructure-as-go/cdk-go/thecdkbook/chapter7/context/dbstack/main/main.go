package main

import (
	"dbstack"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)


func main() {
	app := awscdk.NewApp(nil)

	devVpcId := app.Node().TryGetContext(aws.String("dev-vpc-id")) 
	if devVpcId == nil {
		devVpcId = "vpc-2f09a348"
	}
	prodVpcId := app.Node().TryGetContext(aws.String("prod-vpc-id")) 
	if prodVpcId == nil {
		prodVpcId = "vpc-abcd0123"
	}

	dbstack.NewDbstackStack(app, "DevDb", &dbstack.DbstackStackProps{
		StackProps:   awscdk.StackProps{},
		VpcID:        aws.String(fmt.Sprintf("%v", devVpcId)),
		InstanceType: aws.String("t3.micro"),
	})
	// RDS instance type would be "db.t3.micro", see https://aws.amazon.com/rds/instance-types/
	
	dbstack.NewDbstackStack(app, "ProdDb", &dbstack.DbstackStackProps{
		StackProps:   awscdk.StackProps{},
		VpcID:        aws.String(fmt.Sprintf("%v", prodVpcId)),
		InstanceType: aws.String("r5.xlarge"),
	})
	// RDS instance type would be "db.r5.large", see https://aws.amazon.com/rds/instance-types/
	// EC2 instance types: https://aws.amazon.com/ec2/instance-types/

	app.Synth(nil)
}
