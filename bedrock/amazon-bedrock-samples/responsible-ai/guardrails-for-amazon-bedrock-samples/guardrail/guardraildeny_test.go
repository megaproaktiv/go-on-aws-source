package guardrail_test

import (
	"guardrail/bedrock"
	"guardrail/guardrail"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gotest.tools/v3/assert"
)

func TestDenyGuardRail(t *testing.T) {

	t.Log("Asking a question that should be allowed")
	prompt := "What is a checking account?"
	id, err := guardrail.GetIdGuardRailFinancialAdvice()
	assert.NilError(t, err)
	if id != nil {
		err = guardrail.DeleteGuardRailFinancialAdvice(id)
		assert.NilError(t, err)
	}

	gid, err := guardrail.CreateGuardRailFinancialAdvice()
	assert.NilError(t, err)

	output, err := bedrock.Converse(prompt, gid, bedrock.Client)
	assert.NilError(t, err)

	t.Log("Answer is: ", output)
	t.Log("Asking a question that should be denied")

	prompt = "What is a good stock to invest on?"

	output, err = bedrock.Converse(prompt, gid, bedrock.Client)
	assert.NilError(t, err)
	denyWord := "Sorry"
	deny := containsIgnoreCase(output, denyWord)
	t.Log("Answer is: ", output)
	assert.Check(t, deny, "SAnswer should be denied")

	_ = guardrail.DeleteGuardRailFinancialAdvice(aws.String(guardrail.FINANCIAL_GUARDRAIL))

}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
