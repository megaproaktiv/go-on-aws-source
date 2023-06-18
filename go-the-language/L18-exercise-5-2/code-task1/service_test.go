// begin
package simple_test

import (
	"bytes"
	"gotest.tools/assert"
	"io"
	"os"
	"simple"
	"testing"
)

func TestDo(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()
  // Set the standard output to the pipe
  old := os.Stdout
  os.Stdout = w

  // Call the function here
  
	simple.Do()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Set the standard output to the pipe
	old = os.Stdout
	os.Stdout = w

	assert.Equal(t, "Do\n", buf.String())
	os.Stdout = old

}
