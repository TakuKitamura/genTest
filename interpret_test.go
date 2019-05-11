package main

import (
	"bufio"
	"fmt"
	"genTest/core"
	"os"
	"testing"
)

func TestInterpret(t *testing.T) {
	os.Setenv("GENTEST_MODE", "TEST")

	fp, _ := os.Open("./sample_code/grammar_example.gent")

	scanner := bufio.NewScanner(fp)

	output, _ := core.Exec(scanner)

	expected := `123.3
456
-385786.5333333333255723
163.1333333333333258
Hello, World!
Done!
`
	if output != expected {
		errMsg := fmt.Sprintf("Expected output is \"%s\", but actual is \"%s\".", expected, output)
		t.Fatal(errMsg)
	}
}
