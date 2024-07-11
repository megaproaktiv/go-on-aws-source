package guardrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
)

const FINANCIAL_GUARDRAIL = "financialguardrail"

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
func DeleteGuardRailFinancialAdvice(id *string) error {
	params := &bedrock.DeleteGuardrailInput{
		GuardrailIdentifier: id,
	}
	_, err := Client.DeleteGuardrail(context.TODO(), params)
	if err != nil {
		return err
	}
	return nil

}
