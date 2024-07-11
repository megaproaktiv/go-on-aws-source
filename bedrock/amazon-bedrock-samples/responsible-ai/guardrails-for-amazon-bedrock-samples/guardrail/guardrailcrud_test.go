package guardrail_test

import (
	"guardrail/guardrail"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateGuardCRUD(t *testing.T) {
	t.Logf("CreateGuardRail")
	id, err := guardrail.CreateGuardRailFinancialAdvice()
	assert.NilError(t, err)

	t.Logf("CheckGuardRail")
	storedId, err := guardrail.GetIdGuardRailFinancialAdvice()
	assert.NilError(t, err)
	assert.Equal(t, *id, *storedId, "IDs should be the same")

	t.Logf("DeleteGuardRail")
	err = guardrail.DeleteGuardRailFinancialAdvice(storedId)
	assert.NilError(t, err)

	storedId, err = guardrail.GetIdGuardRailFinancialAdvice()
	assert.NilError(t, err)
	if storedId == nil {
		t.Log("Guardrail has been deleted")
	} else {
		t.Error("Guardrail has not been deleted")
	}

}
