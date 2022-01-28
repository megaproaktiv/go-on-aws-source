package main

import (
	"log"
	"os"
	"instance"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/spf13/viper"
)


func init() {
    
    viper.SetConfigName ("config-dev") 
	viper.AddConfigPath("./config")


    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            // Config file not found; ignore error if desired
            log.Println("no such config file")
        } else {
            // Config file was found but another error was produced
            log.Println("read config error")
        }
    }
}


func main() {
	app := awscdk.NewApp(nil)

	// Hier Nach Account config auswaehlen

	// awscdk.StackProps{
	// 	Env: env(),
	// },

	instance.NewStorageStack(app, "storage", &instance.StorageStackProps{
		StackProps:     awscdk.StackProps{Env: env(),},
		Name:            viper.GetString("instance.name"),
		AMI:            viper.GetString("instance.image"),
		InstanceType:   viper.GetString("instance.type"),
	})


	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {

	return &awscdk.Environment{
	 Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
