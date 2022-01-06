package instance_test

import (
	"log"
	"os"
	"instance"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/aws-sdk-go-v2/aws"
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

func TestServerStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := instance.NewStorageStack(app, "storage", &instance.StorageStackProps{
			StackProps:     awscdk.StackProps{Env: env(),},
			Name:            viper.GetString("instance.name"),
			AMI:            viper.GetString("instance.image"),
			InstanceType:   viper.GetString("instance.type"),
		},
	)
	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(aws.String("AWS::EC2::Instance"), map[string]interface{}{
		"AvailabilityZone": "dummy1a",
	})
}


func env() *awscdk.Environment {

	return &awscdk.Environment{
	 Account: aws.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}