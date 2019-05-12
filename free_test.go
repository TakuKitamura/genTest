package main

import (
	"fmt"
	"genTest/core"
	"os"
	"testing"
)

func TestRPN(t *testing.T) {
	os.Setenv("GENTEST_MODE", "TEST")

	rpnList, err := core.RPN([]byte(`print("ABC" + "DEF")`))
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	object := map[string]interface{}{}
	result, err := core.CalculateByRPN(rpnList, object)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	fmt.Println(string(result), 555)
}
