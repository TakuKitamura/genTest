package main

import (
	"fmt"
	"genTest/core"
	"os"
	"testing"
)

func TestRPN(t *testing.T) {
	os.Setenv("GENTEST_MODE", "TEST")

	rpnList, err := core.RPN([]byte("sqrt(2)*sqrt(2)"))
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	result, err := core.CalculateByRPN(rpnList)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	fmt.Println(string(result), 555)
}
