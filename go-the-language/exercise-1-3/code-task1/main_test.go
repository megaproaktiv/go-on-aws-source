package main_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
    "gotest.tools/assert"

)

func TestMainOutput(t *testing.T) {
    var err error
    cmd := exec.Command("go", "run", "main.go")
    out, err := cmd.CombinedOutput()
    sout := string(out) // because out is []byte
    if err != nil && !strings.Contains(sout, "panic") {
        fmt.Println(sout) // so we can see the full output 
        t.Errorf("%v", err)
    }
    assert.Assert(t,  strings.Contains(sout, "Scope global: 100"))
    assert.Assert(t,  strings.Contains(sout, "Scope main: 200"))   


}

