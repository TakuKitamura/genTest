package main

import (
	"bufio"
	"fmt"
	"genTest/core"
	"os"
)

func main() {
	exitCode := 0

	err := os.Setenv("GENTEST_MODE", "DEVELOP")
	if err != nil {
		fmt.Println(err.Error())
		exitCode = 1
		os.Exit(exitCode)
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: gent file")
		exitCode = 1
		os.Exit(exitCode)
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fp.Close()
		fmt.Println("error: can't open file")
		exitCode = 1
		os.Exit(exitCode)
	}

	scanner := bufio.NewScanner(fp)

	err = core.Exec(scanner)
	if err != nil {
		fmt.Println(err.Error())
		exitCode = 1
		os.Exit(exitCode)
	}
	// fmt.Println(
	// 	"\nOutput",
	// 	"\n---",
	// 	"\n\""+output+"\"",
	// 	"\n---",
	// )
	os.Exit(exitCode)
}
