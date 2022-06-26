package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"table"
)

func main() {
	app := awscdk.NewApp(nil)

	table.NewTableStack(app, "table",nil)

	app.Synth(nil)
}
