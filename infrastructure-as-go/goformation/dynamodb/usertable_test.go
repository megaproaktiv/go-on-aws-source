package phantom_test

import (
	"phantom"
	"testing"

	"gotest.tools/assert"
)



func TestGeneration(t *testing.T){
	content, err := phantom.Synth()
	assert.NilError(t, err, "Generation should work without errors")
	assert.Equal(t, len(*content) > 10, true)
}
