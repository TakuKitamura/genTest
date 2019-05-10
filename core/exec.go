package core

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type StandardFunction string

const (
	Print StandardFunction = "print"
)

const (
	SmallA byte = iota + 97 // a
	SmallB
	SmallC
	SmallD
	SmallE
	SmallF
	SmallG
	SmallH
	SmallI
	SmallJ
	SmallK
	SmallL
	SmallM
	SmallN
	SmallO
	SmallP
	SmallQ
	SmallR
	SmallS
	SmallT
	SmallU
	SmallV
	SmallW
	SmallX
	SmallY
	SmallZ // z
)

const (
	LargeA byte = iota + 65 // A
	LargeB
	LargeC
	LargeD
	LargeE
	LargeF
	LargeG
	LargeH
	LargeI
	LargeJ
	LargeK
	LargeL
	LargeM
	LargeN
	LargeO
	LargeP
	LargeQ
	LargeR
	LargeS
	LargeT
	LargeU
	LargeV
	LargeW
	LargeX
	LargeY
	LargeZ // Z
)

const (
	Zero byte = iota + 48 // 0
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

const (
	NewLine          byte = 10 // \n
	Colon            byte = 58 // :
	DoubleQuotation  byte = 34 // "
	Backslash        byte = 92 // \
	Underscore       byte = 95 // _
	WhiteSpace       byte = 32 //
	Equal            byte = 61 // =
	LeftParenthesis  byte = 40 // (
	RightParenthesis byte = 41 // )
	Comma            byte = 44 // ,
	Dot              byte = 46 // .
	Caret            byte = 94 // ^
	Asterisk         byte = 42 // *
	Slash            byte = 47 // /
	Percent          byte = 37 // %
	Plus             byte = 43 // +
	Minus            byte = 45 // -
)

type CaluculateMap struct {
	Priority int
	F        func(x float64, y float64) float64
}

func caluculateMapFunc() map[string]CaluculateMap {
	caluculateMap := map[string]CaluculateMap{
		string(Caret): CaluculateMap{
			Priority: 2, F: func(x float64, y float64) float64 { return math.Pow(x, y) },
		},
		string(Asterisk): CaluculateMap{
			Priority: 3, F: func(x float64, y float64) float64 { return x * y },
		},
		string(Slash): CaluculateMap{
			Priority: 3, F: func(x float64, y float64) float64 { return x / y },
		},
		string(Percent): CaluculateMap{
			Priority: 3, F: func(x float64, y float64) float64 { return float64(int(x) % int(y)) },
		},
		string(Plus): CaluculateMap{
			Priority: 4, F: func(x float64, y float64) float64 { return x + y },
		},
		string(Minus): CaluculateMap{
			Priority: 4, F: func(x float64, y float64) float64 { return x - y },
		},
	}
	return caluculateMap
}

type Grammar string

const (
	AssignString Grammar = "AssignString"
	Unknown      Grammar = "Unknown"
)

func bytePrint(bytes [][]byte) {
	for _, v := range bytes {
		fmt.Println(string(v) + ", ")
	}
	fmt.Println()
}

func getLastAsciiCode(bytes []byte) byte {
	return bytes[len(bytes)-1]
}

func isAlphabet(char byte) bool {
	if SmallA <= char && char <= SmallZ {
		return true
	}

	if LargeA <= char && char <= LargeZ {
		return true
	}

	return false
}

func isNumber(char byte) bool {
	if Zero <= char && char <= Nine {
		return true
	}

	return false
}

func isMathMark(char byte) bool {
	if char == Caret ||
		char == Asterisk ||
		char == Slash ||
		char == Percent ||
		char == Plus ||
		char == Minus ||
		char == LeftParenthesis ||
		char == RightParenthesis ||
		char == Dot {
		return true
	}
	return false
}

func RPN(formula []byte) ([][]byte, error) {
	stack := [][]byte{}
	rpnList := [][]byte{}
	normalFormulaList := [][]byte{}

	caluculateMap := caluculateMapFunc()

	for {
		matchNumber := regexp.MustCompile(`\d+\.?\d*|\.\d+`).FindSubmatch(formula)
		matchNumberIndex := regexp.MustCompile(`\d+\.?\d*|\.\d+`).FindSubmatchIndex(formula)

		matchOperator := regexp.MustCompile(`[^\d]`).FindSubmatch(formula)

		matchOperatorIndex := regexp.MustCompile(`[^\d]`).FindSubmatchIndex(formula)

		firstMatchNumberIndex := -1
		if len(matchNumberIndex) > 0 {
			firstMatchNumberIndex = matchNumberIndex[0]
		}

		firstmatchOperatorIndex := -1
		if len(matchOperatorIndex) > 0 {
			firstmatchOperatorIndex = matchOperatorIndex[0]
		}
		if len(matchNumber) == 1 && firstMatchNumberIndex == 0 {

			normalFormulaList = append(normalFormulaList, matchNumber[0])
			formula = formula[matchNumberIndex[1]:]
			continue
		}

		if len(matchOperator) == 1 && firstmatchOperatorIndex == 0 {
			normalFormulaList = append(normalFormulaList, matchOperator[0])
			formula = formula[len(matchOperator):]
			continue
		}
		break
	}

	for _, token := range normalFormulaList {
		if bytes.Equal(token, []byte{Dot}) == true {
			errMsg := "unexpected [.]."
			return nil, errors.New(errMsg)
		}

		if regexp.MustCompile(`\d+\.?\d*|\.\d+`).Match(token) == true {
			rpnList = append(rpnList, token)
			continue
		} else if bytes.Equal(token, []byte{RightParenthesis}) == true {
			for {
				if len(stack) == 0 {
					errMsg := "mismatch bracket."
					return nil, errors.New(errMsg)
				}

				topStack := stack[len(stack)-1]
				if bytes.Equal(topStack, []byte{LeftParenthesis}) == true {
					stack = stack[:len(stack)-1]
					break
				} else {
					popedStack := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					rpnList = append(rpnList, popedStack)
				}
			}
			continue
		} else if bytes.Equal(token, []byte{LeftParenthesis}) == true {
			stack = append(stack, token)
			continue
		} else {
			for {
				if len(stack) == 0 {
					stack = append(stack, token)
					break
				} else {
					topStack := stack[len(stack)-1]
					if bytes.Equal(topStack, []byte{LeftParenthesis}) == false {
						_, haveKey := caluculateMap[string(token)]
						if haveKey == false {
							errMsg := "undefined symbol."
							return nil, errors.New(errMsg)
						}

						if caluculateMap[string(token)].Priority < caluculateMap[string(topStack)].Priority {
							stack = append(stack, token)
							break
						} else {
							popedStack := stack[len(stack)-1]
							stack = stack[:len(stack)-1]
							rpnList = append(rpnList, popedStack)
						}
					} else {
						stack = append(stack, token)
						break
					}
				}
			}
		}
	}

	for _, v := range stack {
		if bytes.Equal(v, []byte{LeftParenthesis}) == true || bytes.Equal(v, []byte{RightParenthesis}) == true {
			errMsg := "mismatch bracket."
			return nil, errors.New(errMsg)
		}
		popedStack := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rpnList = append(rpnList, popedStack)
	}

	includeParenthesis := false
	for _, token := range rpnList {
		if bytes.Equal(token, []byte{LeftParenthesis}) == true || bytes.Equal(token, []byte{RightParenthesis}) == true {
			includeParenthesis = true
		}
	}

	if includeParenthesis == true {
		errMsg := "mismatch bracket."
		return nil, errors.New(errMsg)
	}

	if len(rpnList) == 0 {
		errMsg := ""
		return nil, errors.New(errMsg)
	}

	return rpnList, nil
}

func CalculateByRPN(rpnList [][]byte) ([]byte, error) {
	stack := [][]byte{}
	caluculateMap := caluculateMapFunc()

	for _, token := range rpnList {
		existNumber := regexp.MustCompile(`\d+\.?\d*|\.\d+`).Match(token)
		Operaters := [][]byte{[]byte{Caret}, []byte{Asterisk}, []byte{Slash}, []byte{Percent}, []byte{Plus}, []byte{Minus}}
		existOperator := false
		for _, operater := range Operaters {
			if bytes.Equal(token, operater) == true {
				existOperator = true
				break
			}
		}
		if existNumber || existOperator {
			if existNumber {
				stack = append(stack, token)
			} else if existOperator {
				accumulator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				_, haveKey := caluculateMap[string(token)]
				if haveKey == false {
					errMsg := "unexpected token: [' + token + ']."
					return nil, errors.New(errMsg)
				}

				popedStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				x, err := strconv.ParseFloat(string(popedStack), 64)
				if err != nil {
					return nil, err
				}

				y, err := strconv.ParseFloat(string(accumulator), 64)
				if err != nil {
					return nil, err
				}

				accumulator = []byte(fmt.Sprintf("%f", caluculateMap[string(token)].F(x, y)))

				stack = append(stack, accumulator)
			}
		} else {
			if bytes.Equal(token, []byte{LeftParenthesis}) == false && bytes.Equal(token, []byte{RightParenthesis}) == false {
				errMsg := "undefined token: " + string(token)
				return nil, errors.New(errMsg)
			}
		}
	}

	if len(stack) > 1 {
		if bytes.Equal(stack[1], []byte{Dot}) == true {
			errMsg := "unexpected [.]."
			return nil, errors.New(errMsg)
		}
		errMsg := fmt.Sprintf("failed calculate.\nnow stack: %s", stack)
		return nil, errors.New(errMsg)
	}

	return stack[0], nil
}

func Exec(scanner *bufio.Scanner) (string, error) {

	variable := map[string]string{}

	output := ""

	for scanner.Scan() {
		oneLine := scanner.Bytes()
		if len(oneLine) == 0 {
			continue
		}

		lastChar := getLastAsciiCode(oneLine)
		if lastChar == Colon {

		} else {
			inStringState := 0
			inFormulaState := 0
			inAssignState := 0
			inArgumentState := 0
			variableName := []byte{}
			variableValue := []byte{}
			argument := []byte{}

			for i := 0; i < len(oneLine); i++ {
				asciiCode := oneLine[i]

				if inAssignState == 0 {

					if inArgumentState == 0 {
						if isAlphabet(asciiCode) {
							if inStringState == 0 {
								variableName = append(variableName, byte(asciiCode))
								continue
							}
						}

						if asciiCode == LeftParenthesis {
							inArgumentState = 1
							continue
						} else if asciiCode == Equal {
							inAssignState = 1
							continue
						}

					} else if inArgumentState == 1 {
						if asciiCode == RightParenthesis && i == len(oneLine)-1 {
							inArgumentState = 2
							continue
						}
						if asciiCode == DoubleQuotation {
							if inStringState == 0 {
								inStringState = 1
								continue
							} else if inStringState == 1 {
								inStringState = 2
								continue
							}
						} else if inStringState == 0 {
							if isNumber(asciiCode) || isMathMark(asciiCode) {
								argument = append(argument, byte(asciiCode))
								inFormulaState = 1
							} else if isAlphabet(asciiCode) {
								argument = append(argument, byte(asciiCode))
							}
						} else if inStringState == 1 {
							variableValue = append(variableValue, byte(asciiCode))
							// inFormulaState = 1
						}
					} else if inArgumentState == 2 {
						continue
					}
				} else if inAssignState == 1 {
					if asciiCode == DoubleQuotation {
						if inStringState == 0 {
							inStringState = 1
							continue
						} else if inStringState == 1 {
							inStringState = 2
							continue
						}
					} else if isNumber(asciiCode) || isMathMark(asciiCode) {
						variableValue = append(variableValue, byte(asciiCode))
						inFormulaState = 1
					} else {
						if inStringState == 1 {
							variableValue = append(variableValue, byte(asciiCode))
						}
					}

				}

			}

			if inStringState == 1 {
				errMsg := "err: invalid syntax."
				return "", errors.New(errMsg)
			}

			if inArgumentState == 1 {
				errMsg := "err: invalid syntax."
				return "", errors.New(errMsg)
			}
			if os.Getenv("GENTEST_MODE") == "DEVELOP" {
				fmt.Println(
					"\nDebug",
					"\n---",
					"\ninFormulaState: ",
					inFormulaState,
					"\ninAssignState: ",
					inAssignState,
					"\ninArgumentState: ",
					inArgumentState,
					"\ninStringState: ",
					inStringState,
					"\nvariable: ",
					variable,
					"\nvariableName: ",
					string(variableName),
					"\nvariableValue: ",
					string(variableValue),
					"\nargument: ",
					string(argument),
				)
			}

			// a="Hello!"
			if inFormulaState == 0 && inAssignState == 1 && inArgumentState == 0 && inStringState == 2 {
				variable[string(variableName)] = string(variableValue)
				if os.Getenv("GENTEST_MODE") == "DEVELOP" {
					fmt.Println("\nFuncType: 1")
				}
			}

			// a=1+2
			if inFormulaState == 1 && inAssignState == 1 && inArgumentState == 0 && inStringState == 0 {
				if os.Getenv("GENTEST_MODE") == "DEVELOP" {
					fmt.Println("\nFuncType: 2")
				}
				rpnList, err := RPN(variableValue)
				if err != nil {
					return "", err
				}

				result, err := CalculateByRPN(rpnList)
				if err != nil {
					return "", err
				}
				variable[string(variableName)] = string(result)
			}

			// print("Hello!")
			if inFormulaState == 0 && inAssignState == 0 && inArgumentState == 2 && inStringState == 2 {
				if os.Getenv("GENTEST_MODE") == "DEVELOP" {
					fmt.Println("\nFuncType: 3")
				}
				if StandardFunction(variableName) == Print {
					output += fmt.Sprint(string(variableValue)) + "\n"
				}
			}

			// print(msg)
			if inFormulaState == 0 && inAssignState == 0 && inArgumentState == 2 && inStringState == 0 {
				if os.Getenv("GENTEST_MODE") == "DEVELOP" {
					fmt.Println("\nFuncType: 4")
				}
				if StandardFunction(variableName) == Print {
					output += fmt.Sprint(variable[string(argument)]) + "\n"
				}
			}

			// print(1+2)
			if inFormulaState == 1 && inAssignState == 0 && inArgumentState == 2 && inStringState == 0 {
				if os.Getenv("GENTEST_MODE") == "DEVELOP" {
					fmt.Println("\nFuncType: 5")
				}
				if StandardFunction(variableName) == Print {
					rpnList, err := RPN(argument)
					if err != nil {
						return "", err
					}

					result, err := CalculateByRPN(rpnList)
					if err != nil {
						return "", err
					}
					output += fmt.Sprint(string(result)) + "\n"
				}
			}

			if os.Getenv("GENTEST_MODE") == "DEVELOP" {
				fmt.Println(
					"\nnowOutput:",
					"\n\""+output+"\"",
					"\n---",
				)
			}

		}
	}

	return output, nil
}
