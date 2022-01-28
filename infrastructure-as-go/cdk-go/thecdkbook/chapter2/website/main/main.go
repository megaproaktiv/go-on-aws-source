package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	website "MySimpleWebsite"
)


func main() {
	app := awscdk.NewApp(nil)

	website.NewWebsiteStack(app, "SomeWebsitecom", nil)

	app.Synth(nil)
}
