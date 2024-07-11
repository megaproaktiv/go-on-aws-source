package bedrock_test

import (
	"context"
	"encoding/json"
	"guardrail/bedrock"
	"guardrail/guardrail"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"gotest.tools/v3/assert"
)

func TestApplyGuard(t *testing.T) {

	id, err := guardrail.GetIdGuardRailFinancialAdvice()
	err = guardrail.DeleteGuardRailFinancialAdvice(id)
	time.Sleep(5 * time.Second)
	gid, err := guardrail.CreateGuardRailFinancialAdvice()
	assert.NilError(t, err)

	client := bedrock.Client
	prompt := aws.String("What is a checking account?")
	params := &bedrockruntime.ApplyGuardrailInput{
		Content: []types.GuardrailContentBlock{
			&types.GuardrailContentBlockMemberText{
				Value: types.GuardrailTextBlock{
					Text: prompt,
					Qualifiers: []types.GuardrailContentQualifier{
						types.GuardrailContentQualifierGuardContent,
					},
				},
			},
		},
		GuardrailIdentifier: gid,
		GuardrailVersion:    aws.String("DRAFT"),
		Source:              types.GuardrailContentSourceInput,
	}
	resp, err := client.ApplyGuardrail(context.TODO(), params)

	assert.NilError(t, err)
	// Pretty-print the JSON response
	respJson, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal response to JSON: %v", err)
	}

	t.Logf("Response: %s", string(respJson))
	guardrail.DeleteGuardRailFinancialAdvice(aws.String(guardrail.FINANCIAL_GUARDRAIL))

}
