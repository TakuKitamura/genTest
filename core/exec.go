package core

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

type StandardGrammer string

const (
	If StandardGrammer = "if"
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
	NewLine           byte = 10  // \n
	Colon             byte = 58  // :
	DoubleQuotation   byte = 34  // "
	Backslash         byte = 92  // \
	Underscore        byte = 95  // _
	WhiteSpace        byte = 32  //
	Equal             byte = 61  // =
	LeftParenthesis   byte = 40  // (
	RightParenthesis  byte = 41  // )
	Comma             byte = 44  // ,
	Dot               byte = 46  // .
	Caret             byte = 94  // ^
	Asterisk          byte = 42  // *
	Slash             byte = 47  // /
	Percent           byte = 37  // %
	Plus              byte = 43  // +
	Minus             byte = 45  // -
	ExclamationMark   byte = 33  // !
	LeftAngleBracket  byte = 60  // <
	RightAngleBracket byte = 62  // >
	Ampersand         byte = 38  // &
	Pipeline          byte = 124 // |
)

type OperatorMap struct {
	OperatorType           OperaterType
	Priority               int
	ArithmeticFunc         func(x interface{}, y interface{}) (interface{}, error)
	EqualityComparisonFunc func(x interface{}, y interface{}) (bool, error)
}

func p(input interface{}) {
	//	fmt.Println("debugLog:", input)
}

type StandardFunctionName string

const (
	Print StandardFunctionName = "print"
)

type OperaterType string

const (
	Arithmetic         OperaterType = "Arithmetic"
	EqualityComparison OperaterType = "EqualityComparison"
	Function           OperaterType = "Function"
)

type StandardFunctionMap struct {
	OperatorType OperaterType
	Priority     int
	Argc         int
	NormalFunc   func(x ...interface{}) (interface{}, error)
}

func standardFunctionMap() map[string]StandardFunctionMap {
	standardFunctionMap := map[string]StandardFunctionMap{
		string("sqrt"): StandardFunctionMap{
			OperatorType: Function,
			Argc:         1,
			Priority:     1,
			NormalFunc: func(x ...interface{}) (interface{}, error) {
				if len(x) != 1 {
					errMsg := fmt.Sprintf("%s args is invalid.", string("sqrt"))
					return nil, errors.New(errMsg)
				}

				if xv, ok := x[0].(float64); ok {
					sqrtValue := interface{}(math.Sqrt(xv))
					return sqrtValue, nil
				}

				errMsg := fmt.Sprintf("can't cast with %s", string("sqrt"))
				return nil, errors.New(errMsg)
			},
		},
		string("print"): StandardFunctionMap{
			OperatorType: Function,
			Argc:         1,
			Priority:     1,
			NormalFunc: func(x ...interface{}) (interface{}, error) {
				output := "("
				for i, v := range x {
					if xv, ok := v.([]byte); ok {
						xv = bytes.Replace(xv, []byte("\\n"), []byte{NewLine}, -1)
						output += string(xv)
					} else {
						errMsg := fmt.Sprintf("can't cast with %s", string("print"))
						return false, errors.New(errMsg)
					}

					if i != len(x)-1 {
						output += ", "
					}
				}
				//	fmt.Println(9999)
				output += ")"
				fmt.Println(output)
				return nil, nil
			},
		},
	}
	return standardFunctionMap
}

func operatorMap() map[string]OperatorMap {
	operatorMap := map[string]OperatorMap{

		string(Asterisk): OperatorMap{
			OperatorType: Arithmetic,
			Priority:     3,
			ArithmeticFunc: func(x interface{}, y interface{}) (interface{}, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return interface{}(xv * yv), nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(Asterisk))
				return false, errors.New(errMsg)
			},
		},
		string(Slash): OperatorMap{
			OperatorType: Arithmetic,
			Priority:     3,
			ArithmeticFunc: func(x interface{}, y interface{}) (interface{}, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return interface{}(xv / yv), nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(Slash))
				return false, errors.New(errMsg)
			},
		},
		string(Percent): OperatorMap{
			OperatorType: Arithmetic,
			Priority:     3,

			ArithmeticFunc: func(x interface{}, y interface{}) (interface{}, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return interface{}(int(xv) % int(yv)), nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(Percent))
				return false, errors.New(errMsg)
			},
		},
		string(Plus): OperatorMap{
			OperatorType: Arithmetic,
			Priority:     4,
			ArithmeticFunc: func(x interface{}, y interface{}) (interface{}, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return interface{}(xv + yv), nil
					}
				}

				if xv, ok := x.(string); ok {
					if yv, ok := y.(string); ok {
						return interface{}(xv + yv), nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(Plus))
				return false, errors.New(errMsg)
			},
		},
		string(Minus): OperatorMap{
			OperatorType: Arithmetic,
			Priority:     4,
			ArithmeticFunc: func(x interface{}, y interface{}) (interface{}, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return interface{}(xv - yv), nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(Minus))
				return false, errors.New(errMsg)
			},
		},
		string(LeftAngleBracket): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     6,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return xv < yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(LeftAngleBracket))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{LeftAngleBracket, Equal}): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     6,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return xv > yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string([]byte{LeftAngleBracket, Equal}))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{RightAngleBracket, Equal}): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     6,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return xv >= yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string([]byte{RightAngleBracket, Equal}))
				return false, errors.New(errMsg)
			},
		},
		string(RightAngleBracket): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     6,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return xv > yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string(RightAngleBracket))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{Equal, Equal}): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     7,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return xv == yv, nil
					}
				}

				if xv, ok := x.(string); ok {
					if yv, ok := y.(string); ok {
						return xv == yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string([]byte{Equal, Equal}))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{ExclamationMark, Equal}): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     7,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(float64); ok {
					if yv, ok := y.(float64); ok {
						return xv != yv, nil
					}
				}

				if xv, ok := x.(string); ok {
					if yv, ok := y.(string); ok {
						return xv != yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string([]byte{ExclamationMark, Equal}))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{Ampersand, Ampersand}): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     11,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(bool); ok {
					if yv, ok := y.(bool); ok {
						return xv && yv, nil
					}
				}

				errMsg := fmt.Sprintf("can't cast with %s", string([]byte{Ampersand, Ampersand}))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{Pipeline, Pipeline}): OperatorMap{
			OperatorType: EqualityComparison,
			Priority:     12,
			EqualityComparisonFunc: func(x interface{}, y interface{}) (bool, error) {
				if xv, ok := x.(bool); ok {
					if yv, ok := y.(bool); ok {
						return xv || yv, nil
					}
				}
				errMsg := fmt.Sprintf("can't cast with %s", string([]byte{Ampersand, Ampersand}))
				return false, errors.New(errMsg)
			},
		},
		string([]byte{Equal}): OperatorMap{
			Priority: 14,
		},
		string([]byte{Comma}): OperatorMap{
			Priority: 15,
		},
	}
	return operatorMap
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

	operatorsAndParenthesis := [][]byte{
		[]byte{LeftParenthesis},
		[]byte{RightParenthesis},
	}

	operatorMap := operatorMap()

	for key, _ := range operatorMap {
		operatorsAndParenthesis = append(operatorsAndParenthesis, []byte(key))
	}

	sort.Slice(operatorsAndParenthesis, func(i, j int) bool { return len(operatorsAndParenthesis[i]) > len(operatorsAndParenthesis[j]) })

	inStringSentence := 0
	tempFormula := []byte{}
	for i := 0; i < len(formula); i++ {

		if inStringSentence == 0 {
			if formula[i] == WhiteSpace {
				continue
			}
		}

		if formula[i] == DoubleQuotation {
			if inStringSentence == 0 {
				inStringSentence = 1
			} else if inStringSentence == 1 {
				inStringSentence = 0
			}
		}

		tempFormula = append(tempFormula, formula[i])
	}
	//	fmt.Println(string(tempFormula), 777)
	formula = tempFormula

	for {
		//	fmt.Println(string(formula))
		matchNumberIndex := regexp.MustCompile(`^\d+\.?\d*|\.\d+`).FindIndex(formula)
		matchNumber := []byte{}
		if len(matchNumberIndex) > 0 {
			if matchNumberIndex[0] == 0 {
				matchNumber = formula[matchNumberIndex[0]:matchNumberIndex[1]]

				normalFormulaList = append(normalFormulaList, matchNumber)
				formula = formula[matchNumberIndex[1]:]
				continue
			}
		}

		matchAlphabetIndex := regexp.MustCompile(`^[A-Za-z]+`).FindIndex(formula)
		matchAlphabet := []byte{}
		if len(matchAlphabetIndex) > 0 {
			if matchAlphabetIndex[0] == 0 {
				matchAlphabet = formula[matchAlphabetIndex[0]:matchAlphabetIndex[1]]

				normalFormulaList = append(normalFormulaList, matchAlphabet)
				formula = formula[matchAlphabetIndex[1]:]
				continue
			}
		}

		matchStringIndex := regexp.MustCompile(`^".*?"`).FindIndex(formula)
		matchString := []byte{}
		if len(matchStringIndex) > 0 {
			if matchStringIndex[0] == 0 {
				matchString = formula[matchStringIndex[0]:matchStringIndex[1]]

				normalFormulaList = append(normalFormulaList, matchString)
				formula = formula[matchStringIndex[1]:]
				continue
			}
		}

		isMatchOperator := false
		for _, v := range operatorsAndParenthesis {
			if bytes.HasPrefix(formula, v) == true {
				matchOperator := formula[:len(v)]
				normalFormulaList = append(normalFormulaList, matchOperator)
				formula = formula[len(matchOperator):]
				isMatchOperator = true
				break
			}
		}

		if isMatchOperator == true {
			continue
		}

		break
	}
	// fmt.Println("222")
	// bytePrint(normalFormulaList)
	// fmt.Println("111")

	for _, token := range normalFormulaList {
		if bytes.Equal(token, []byte{Dot}) == true {
			errMsg := "unexpected [.]."
			return nil, errors.New(errMsg)
		}

		// if bytes.Equal(token, []byte{Comma}) == true {
		//	fmt.Println("TOKEN: ", string(token))
		//	fmt.Println("RPNLIST: ")
		// bytePrint(rpnList)
		//	fmt.Println("STACK: ")
		// bytePrint(stack)
		//	fmt.Println()

		// continue
		// }

		if bytes.Equal(token, []byte{Comma}) == true {

		}

		if regexp.MustCompile(`^\d+\.?\d*|\.\d+`).Match(token) == true {
			rpnList = append(rpnList, token)
			continue
		} else if regexp.MustCompile(`^[A-Za-z]+`).Match(token) == true {
			stack = append(stack, token)
			continue
		} else if regexp.MustCompile(`^".*?"`).Match(token) == true {
			stack = append(stack, token)
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
						_, haveKey := operatorMap[string(token)]
						// fmt.Println(string(token) + "|")
						if haveKey == false {
							errMsg := "undefined symbol."
							return nil, errors.New(errMsg)
						}

						//	fmt.Println("AAAAAA: ", string(token), operatorMap[string(token)].Priority)
						if operatorMap[string(token)].Priority < operatorMap[string(topStack)].Priority {
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
		errMsg := "rpnList is empty."
		return nil, errors.New(errMsg)
	}

	// bytePrint(rpnList)
	//	fmt.Println(99999)

	return rpnList, nil
}

func CalculateByRPN(rpnList [][]byte, object map[string][]byte) error {
	// bytePrint(rpnList)
	stack := [][]byte{}

	operatorMap := operatorMap()
	operators := [][]byte{}

	for key, _ := range operatorMap {
		operators = append(operators, []byte(key))
	}

	// sort.Slice(operators, func(i, j int) bool { return len(operators[i]) > len(operators[j]) })
	for _, token := range rpnList {
		existNumber := regexp.MustCompile(`^\d+\.?\d*|\.\d+`).Match(token)
		existAlphabet := regexp.MustCompile(`^[A-Za-z]+`).Match(token)
		existString := regexp.MustCompile(`^".*?"`).Match(token)

		// Operaters := [][]byte{[]byte{Equal, Equal}, []byte{Caret}, []byte{Asterisk}, []byte{Slash}, []byte{Percent}, []byte{Plus}, []byte{Minus}}
		existOperator := false
		for _, operator := range operators {
			if bytes.Equal(token, operator) == true {
				existOperator = true
				break
			}
		}

		standardFunctionMap := standardFunctionMap()
		standardFunctions := [][]byte{}

		existStandardFunction := false

		for key, _ := range standardFunctionMap {
			standardFunctions = append(standardFunctions, []byte(key))
		}

		for _, standardFunction := range standardFunctions {
			if bytes.Equal(token, standardFunction) == true {
				existStandardFunction = true
				break
			}
		}

		//	fmt.Println("token, ", string(token))
		//	fmt.Println("nowStack1, ")
		// bytePrint(stack)
		//	fmt.Println()

		if existNumber {
			stack = append(stack, token)
		} else if existOperator {
			if bytes.Equal(token, []byte{Equal}) == true {
				//	fmt.Println(777938475879234)
				// bytePrint((stack))
				object[string(stack[0])] = stack[1]
				stack = [][]byte{}
			} else if bytes.Equal(token, []byte{Comma}) == true {
				// stack = append(stack, token)
				//	fmt.Println("hello!")
			} else {
				accumulator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				_, haveKey := operatorMap[string(token)]
				if haveKey == false {
					errMsg := "unexpected token: [' + token + ']."
					return errors.New(errMsg)
				}

				popedStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//	fmt.Println(123)

				var x interface{}
				var y interface{}
				var err error

				for {

					if regexp.MustCompile(`^\d+\.?\d*|\.\d+`).Match(popedStack) == true && regexp.MustCompile(`^\d+\.?\d*|\.\d+`).Match(accumulator) {
						x, err = strconv.ParseFloat(string(popedStack), 64)
						if err != nil {
							return err
						}

						y, err = strconv.ParseFloat(string(accumulator), 64)
						if err != nil {
							return err
						}
						break
					} else if regexp.MustCompile(`^".*?"`).Match(popedStack) == true && regexp.MustCompile(`^".*?"`).Match(accumulator) {

						x = string(popedStack[:len(popedStack)-1])

						y = string(accumulator[1:])
						break

					} else if (string(popedStack) == "true" || string(popedStack) == "false") || (string(accumulator) == "true" || string(accumulator) == "false") {
						x, err = strconv.ParseBool(string(popedStack))
						if err != nil {
							return err
						}
						y, err = strconv.ParseBool(string(accumulator))
						if err != nil {
							return err
						}
						break
					} else {

						// x = string(popedStack[:len(popedStack)-1])

						xv, haveKey := object[string(popedStack)]
						if haveKey == true {
							popedStack = xv
						}

						// y = string(accumulator[1:])
						//	fmt.Println(x, y, 123456)

						yv, haveKey := object[string(accumulator)]
						if haveKey == true {
							accumulator = yv
						}

						fmt.Println(popedStack, accumulator, xv, yv)

						if len(xv) == 0 && len(yv) == 0 {
							errMsg := "x or y is invalid."
							return errors.New(errMsg)
						}
					}

				}

				//	fmt.Println(456)

				// operatorMap := operatorMap[string(token)]
				operatorMap, haveKey := operatorMap[string(token)]
				if haveKey == true {
					if operatorMap.OperatorType == Arithmetic {
						v, err := operatorMap.ArithmeticFunc(x, y)
						if err != nil {
							return err
						}

						if vFloat64, ok := v.(float64); ok {
							accumulator = []byte(strconv.FormatFloat(vFloat64, 'f', 16, 64))
						} else if vString, ok := v.(string); ok {
							accumulator = []byte(vString)
						} else {
							errMsg := "unknown InterfaceType"
							return errors.New(errMsg)
						}

					} else if operatorMap.OperatorType == EqualityComparison {
						vBool, err := operatorMap.EqualityComparisonFunc(x, y)

						if err != nil {
							return err

						}
						accumulator = []byte(strconv.FormatBool(vBool))
					} else {
						errMsg := "undefined OperatorType: " + string(operatorMap.OperatorType)
						return errors.New(errMsg)
					}
				} else {
					errMsg := "don't have token: " + string(token)
					return errors.New(errMsg)
				}

				stack = append(stack, accumulator)

			}
		} else if existStandardFunction {

			if len(stack) == 0 {
				errMsg := fmt.Sprintf("%s mistake usage.", string(token))
				return errors.New(errMsg)
			}

			_, haveKey := standardFunctionMap[string(token)]
			if haveKey == false {
				errMsg := "unexpected token: [' + token + ']."
				return errors.New(errMsg)
			}

			argc := standardFunctionMap[string(token)].Argc

			if string(token) == "sqrt" {
				sqrtArgs := []interface{}{}
				//	fmt.Println(argc, 123)

				for i := len(stack) - 1; i >= len(stack)-argc; i-- {

					var f64 float64
					var err error
					v, haveKey := object[string(stack[i])]
					if haveKey == true {
						f64, err = strconv.ParseFloat(string(v), 64)
						if err != nil {
							return err
						}
					} else {
						f64, err = strconv.ParseFloat(string(stack[i]), 64)
						if err != nil {
							return err
						}
					}

					sqrtArgs = append(sqrtArgs, f64)
				}

				v, err := standardFunctionMap[string(token)].NormalFunc(sqrtArgs...)
				if err != nil {
					return err
				}

				if vFloat64, ok := v.(float64); ok {
					answer := []byte(strconv.FormatFloat(vFloat64, 'f', 16, 64))
					stack = stack[:len(stack)-argc]
					stack = append(stack, answer)
				} else {
					errMsg := "unknown InterfaceType"
					return errors.New(errMsg)
				}
			} else if string(token) == "print" {
				printArgs := []interface{}{}
				//	fmt.Println(argc, 123)
				// for i := len(stack) - 1; i >= len(stack)-argc; i-- {
				// 	// f64, err := strconv.ParseFloat(string(stack[i]), 64)
				// 	// if err != nil {
				// 	// 	return err
				// 	// }
				// 	printArgs = append(printArgs, stack[i])
				// }

				// for i := 0; i < argc; i++ {
				for i := 0; i < len(stack); i++ {
					// f64, err := strconv.ParseFloat(string(stack[i]), 64)
					// if err != nil {
					// 	return err
					// }
					//	fmt.Println(object, 1212121212)
					existVariable := regexp.MustCompile(`^[A-Za-z]+`).Match(stack[i])
					if string(stack[i]) == "true" || string(stack[i]) == "false" {
						//	fmt.Println(7823873287232783)
						printArgs = append(printArgs, stack[i])
						continue
					}

					if existVariable == true {
						printArgs = append(printArgs, object[string(stack[i])])
						continue
					}

					printArgs = append(printArgs, stack[i])

				}
				standardFunctionMap[string(token)].NormalFunc(printArgs...)

				// stack = stack[:len(stack)-argc]

				//	fmt.Println(88888)
				// if err != nil {
				// 	return err
				// }

				stack = [][]byte{}

				// if vFloat64, ok := v.(float64); ok {
				// 	answer := []byte(strconv.FormatFloat(vFloat64, 'f', 16, 64))
				// 	stack = stack[:len(stack)-argc]
				// 	stack = append(stack, answer)
				// } else {
				// 	errMsg := "unknown InterfaceType"
				// 	return errors.New(errMsg)
				// }
			}

		} else if existAlphabet {
			stack = append(stack, token)
		} else if existString {
			stack = append(stack, token)
		} else {
			if bytes.Equal(token, []byte{LeftParenthesis}) == false && bytes.Equal(token, []byte{RightParenthesis}) == false {
				errMsg := "undefined token: " + string(token)
				return errors.New(errMsg)
			}
		}

		//	fmt.Println("nowStack2, ")
		// bytePrint(stack)
		//	fmt.Println("---")
		//	fmt.Println()

	}

	//	fmt.Println(object)

	// bytePrint(stack)
	// fmt.Println("here")

	// for {
	// 	popedStack := stack[len(stack)-1]
	// //	fmt.Println("popedStack,", string(popedStack))
	// 	stack = stack[:len(stack)-1]
	// 	if len(stack) == 0 {
	// 		break
	// 	}
	// }
	if len(stack) > 1 {
		if bytes.Equal(stack[1], []byte{Dot}) == true {
			errMsg := "unexpected [.]."
			return errors.New(errMsg)
		}
		errMsg := fmt.Sprintf("failed calculate.\nnow stack: %s", stack)
		return errors.New(errMsg)
	}

	// output := []byte{}

	// if len(stack) > 0 {
	// 	output = stack[0]
	// } else {
	// 	output = nil
	// }

	return nil
}

func Exec(scanner *bufio.Scanner) error {

	object := map[string][]byte{}

	for scanner.Scan() {
		oneLine := scanner.Bytes()
		if len(oneLine) == 0 {
			continue
		}

		rpnList, err := RPN(oneLine)
		if err != nil {
			return err
		}

		err = CalculateByRPN(rpnList, object)
		if err != nil {
			return err
		}
	}
	return nil
}
