//build cloud local
package lambdasummary

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
)

//go:generate moq -out assume_moq_test.go . AssumeInterface

// AssumeInterface Interface for sts
type AssumeInterface interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

// GetCfgSub assume role in sub account
func GetCfgSub(client AssumeInterface, member string) (aws.Config, error) {

	roleArn := "arn:aws:iam::" + member + ":role/CrossAccountListFunctionsRole"
	sessionname := "showfunctions"
	params := &sts.AssumeRoleInput{
		RoleArn:         &roleArn,
		RoleSessionName: &sessionname,
	}
	credentialsSubResponse, err := client.AssumeRole(context.TODO(), params)

	if err != nil {
		log.Error(err)
		return aws.Config{}, err
	}
	credentialsSub := *credentialsSubResponse.Credentials

	cfgSub, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     *credentialsSub.AccessKeyId,
				SecretAccessKey: *credentialsSub.SecretAccessKey,
				SessionToken:    *credentialsSub.SessionToken,
				Source:          "assumerole",
			},
		}))
	if err != nil {
		log.Error(err)
		return aws.Config{}, err
	}

	return cfgSub, nil
}
