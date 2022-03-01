package main

import (
    "cdk.tf/go/stack/generated/hashicorp/aws"
    "cdk.tf/go/stack/generated/hashicorp/aws/ec2"

    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
    "github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
    stack := cdktf.NewTerraformStack(scope, &id)

    aws.NewAwsProvider(stack, jsii.String("aws"), &aws.AwsProviderConfig{
        Region: jsii.String("eu-central-1"),
    })

    instance := ec2.NewInstance(stack, jsii.String("cdktfgo"), &ec2.InstanceConfig{
        Ami:          jsii.String("ami-04c921614424b07cd"),
        InstanceType: jsii.String("t2.micro"),
    })

    cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
        Value: instance.PublicIp(),
    })

    return stack
}

func main() {
    app := cdktf.NewApp(nil)

    NewMyStack(app, "learn-cdktf-go")

    app.Synth()
}

