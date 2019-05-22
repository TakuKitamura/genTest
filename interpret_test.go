package main

import (
	"bufio"
	"bytes"
	"fmt"
	"genTest/core"
	"os"
	"testing"
)

func TestInterpret(t *testing.T) {
	os.Setenv("GENTEST_MODE", "TEST")

	fp, _ := os.Open("./sample_code/grammar_example.gent")

	scanner := bufio.NewScanner(fp)

	w := new(bytes.Buffer)

	err := core.Exec(scanner, w)

	if err != nil {
		errMsg := err.Error()
		t.Fatal(errMsg)
	}

	output := w.String()

	expected := `このようなことしかできません｡
できることその1: 演算
sqrt((2.0 + (3.0 - 4.0) / 5.0) + 0.1 * 2.0) is 1.4142135623730951
できることその2: 評価
1 < 2 && (3 != 3 || true) is true
できることその3: 文字列結合
できることその4: if文による分岐
2020 年は
うるう年です｡
`

	if output != expected {
		errMsg := fmt.Sprintf("Expected output is \"%s\", but actual is \"%s\".", expected, output)
		t.Fatal(errMsg)
	}
}
