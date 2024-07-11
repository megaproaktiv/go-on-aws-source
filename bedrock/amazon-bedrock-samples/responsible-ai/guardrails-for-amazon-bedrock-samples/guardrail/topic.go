package guardrail

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
)

const FINANCIAL_GUARDRAIL = "financialguardrail"
const VERSION = "DRAFT"

// Create a finnancial guardrail
func CreateGuardRailFinancialAdvice() (*string, error) {
	params := &bedrock.CreateGuardrailInput{
		Name: aws.String(FINANCIAL_GUARDRAIL),
		TopicPolicyConfig: &types.GuardrailTopicPolicyConfig{
			TopicsConfig: []types.GuardrailTopicConfig{
				{
					Name:       aws.String("FinancialAdvise"),
					Definition: aws.String("Anything related to provide financial advise, investment recommendations, or similar."),
					Examples: []string{
						"Should I buy this stock?",
						"Should I invest in AMZN stock?",
						"Whats included in my tax declaration?",
					},
					Type: types.GuardrailTopicTypeDeny,
				},
			},
		},
		BlockedInputMessaging:   aws.String("Sorry I cannot respond to that."),
		BlockedOutputsMessaging: aws.String("Sorry I cannot respond to that."),
	}
	ctx := context.TODO()
	resp, err := Client.CreateGuardrail(ctx, params)
	if err != nil {
		return nil, err
	}
	return resp.GuardrailId, nil
}

// Get ID of the created guardrail
func GetIdGuardRailFinancialAdvice() (*string, error) {

	var id *string
	resp, err := Client.ListGuardrails(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	// New GO 1.22 range
	for i := range resp.Guardrails {
		if (*resp.Guardrails[i].Name == FINANCIAL_GUARDRAIL) && (*&resp.Guardrails[i].Status != "DELETING") {
			id = resp.Guardrails[i].Id
		}
	}

	return id, nil

}

// Delete the guardrail
// To delete a guardrail, only specify the ARN of the guardrail in the guardrailIdentifier field.
// If you delete a guardrail, all of its versions will be deleted.
func DeleteGuardRailFinancialAdvice(id *string) error {
	params := &bedrock.DeleteGuardrailInput{
		GuardrailIdentifier: id,
	}
	_, err := Client.DeleteGuardrail(context.TODO(), params)
	if err != nil {
		var notFoundErr *types.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
			fmt.Println("Guardrail not found, nothing to delete.")
			return nil
		}
		return err
	}

	maxCount := 10
	count := 0
	for {
		resp, err := Client.GetGuardrail(context.TODO(), &bedrock.GetGuardrailInput{
			GuardrailIdentifier: id,
		})
		if err != nil {
			var notFoundErr *types.ResourceNotFoundException
			if errors.As(err, &notFoundErr) {
				fmt.Println("Guardrail is deleted.")
				return nil
			} else {
				fmt.Println("Error getting guardrail status")
				return err
			}
		}
		if resp.Status == "DELETING" {
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		}
		count++
		if count > maxCount {
			return errors.New("Timeout waiting for guardrail to be deleted")

		}
	}

	return nil

}
