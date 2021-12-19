package dsl_test

import (
	"os"
	"testing"
	"gotest.tools/assert"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/megaproaktiv/cit/citlambda"
	
)

func TestInfraLambdaExists(t *testing.T) {
	if os.Getenv("I_TEST") != "yes" {
		t.Skip("Skipping testing in non Integration environment")
	}
	gotFc, err := citlambda.GetFunctionConfiguration(aws.String("dsl"),aws.String("myHandler"))
	assert.NilError(t, err, "GetFunctionConfiguration should return no error")
	
	expectHandler := "main"
	assert.Equal(t, expectHandler, *gotFc.Handler )
}