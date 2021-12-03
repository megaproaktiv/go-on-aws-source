package main

import (
	"vpc"
	"github.com/aws/aws-cdk-go/awscdk/v2"

)

func main() {
	app := awscdk.NewApp(nil)

	vpc.VpcStack(app, "vpc", nil)

	app.Synth(nil)
}

