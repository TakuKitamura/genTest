package main

import (
	"fmt"
	"genTest/core"
	"os"
	"testing"
)

func TestRPN(t *testing.T) {
	os.Setenv("GENTEST_MODE", "TEST")

	rpnList, err := core.RPN([]byte(`a=sqrt(sqrt(9))*2`))
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	object := map[string][]byte{}
	err = core.CalculateByRPN(rpnList, object)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	fmt.Println(object, 777)
}
