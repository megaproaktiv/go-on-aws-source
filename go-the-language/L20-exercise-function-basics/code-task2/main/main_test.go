package main_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
    //begin import
    "gotest.tools/assert"
    //end import

)

//begin test
func TestMainOutput(t *testing.T) {
    var err error
    cmd := exec.Command("go", "run", "main.go")
    out, err := cmd.CombinedOutput()
    sout := string(out) // because out is []byte
    if err != nil && !strings.Contains(sout, "panic") {
        fmt.Println(sout) // so we can see the full output 
        t.Errorf("%v", err)
    }
//end test
    expect := `Walk into a bar story:
Bob and Alice walked into a bar.
Bob is excited.
Alice is relaxed.
Bartender is curious.
Bartender said to [Alice Bob], "What brings you in tonight?"
Alice is excited.
Bob is excited.
Bob said to [Alice], "Will you marry me?"
Alice is excited.
Alice said to [Bob], "Yes! I will!"
Bob is happy.
Bartender said to [Bob Alice], "Congratulations to both of you! First round's on me."
Bob is grateful.
Alice is grateful.
End
`   //begin test
    // set expect to expected string (to long to include)
    assert.Equal(t,  expect, sout, "Expected output")
    //end test
    


}

