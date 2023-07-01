package assume

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
)

var Client *sts.Client

// Initialize the client
func init() {
	// Create the client
	//begin client1
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = sts.NewFromConfig(cfg)
	//end client1
}

// GetCfgSub assume role in sub account
func GetCfgSub(client *sts.Client, roleArn *string) (*aws.Config, error) {

	//begin assume
	params := &sts.AssumeRoleInput{
		RoleArn:         roleArn,
		RoleSessionName: aws.String("SecondLife"),
	}
	credentialsSubResponse, err := client.AssumeRole(context.TODO(), params)
	//end assume

	if err != nil {
		log.Println(err)
		return nil, err
	}

	//begin credentials
	credentialsSub := *credentialsSubResponse.Credentials
	cfgSub, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     *credentialsSub.AccessKeyId,
				SecretAccessKey: *credentialsSub.SecretAccessKey,
				SessionToken:    *credentialsSub.SessionToken,
				Source:          "assumerole",
			},
		}))
	//end credentials

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &cfgSub, nil
}
