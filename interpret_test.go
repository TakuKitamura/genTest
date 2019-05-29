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
(2.0 + (3.0 - 4.0) / 5.0) + 0.1 * 2.0 is 2.0000000000000000
できることその2: 評価
1 < 2 && (3 != 3 || true) is true
できることその3: 文字列結合
できることその4: IF文による分岐
2020 年は
うるう年です｡
できることその5: FOR文による分岐
ライプニッツの公式(n=10)より
π ≒ 3.0418396189294028
`

	if output != expected {
		errMsg := fmt.Sprintf("Expected output is \"%s\", but actual is \"%s\".", expected, output)
		t.Fatal(errMsg)
	}
}
