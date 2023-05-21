package table

import (
	"testing"

	"gotest.tools/assert"
)

func TestReadFile(t *testing.T) {

	got, err := ReadFile()
	assert.NilError(t, err)
	assert.Equal(t, len(got), 2)
	assert.Equal(t, got[0].ISBN, "111-1111111111")
	
	
}
