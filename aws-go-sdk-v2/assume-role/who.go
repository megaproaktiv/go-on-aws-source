package assume

import (
	"context"
	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Show the current user with sts get-caller-identity
func Who(client *sts.Client) (*string, error) {
	// Call get-caller-identity
	resp, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	// Handle errors
	if err != nil {
		return nil, err
	}
	// Return the user
	return resp.UserId, nil
}

// Show the account
func Where(client *sts.Client) (*string, error) {
	// Call get-caller-identity
	resp, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	// Handle errors
	if err != nil {
		return nil, err
	}
	// Return the user
	return resp.Account, nil
}
