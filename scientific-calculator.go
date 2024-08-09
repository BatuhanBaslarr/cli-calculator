package main

import "fmt"

func sciCalc() {
	var operation, trigonometry string
	var num1, num2, result float64

	for {
		if result == 0 {
			fmt.Println("\nEnter 1st Number: ")
			num1 = scanWithTypeCheck()

		} else {
			num1 = result
		}

		for {
			fmt.Print("\nEnter operation (+, -, *, /, sqrt, pow, sqr, log, ln, fact, abs, trigo):\n ")
			fmt.Scan(&operation)

			switch operation {
			case "+":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
				result = addVal(num1, num2)
			case "-":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
				result = subVal(num1, num2)
			case "*":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
				result = mulVal(num1, num2)
			case "/":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
				result = divVal(num1, num2)
			case "sqrt":
				result = sqrtVal(num1)
			case "pow":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
				result = powVal(num1, num2)
			case "sqr":
				result = sqrVal(num1)
			case "log":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
				if num2 <= 0 || num2 == 1 {
					fmt.Println("\nLogarithm base must be greater than zero and not equal to one")
					continue
				}
				result = logVal(num1, num2)
				graphFunc(num1, num2, operation)
			case "ln":
				if num1 <= 0 {
					fmt.Println("\nLn logarithm argument must be greater than zero")
					fmt.Println("\nEnter 1st number:")
					fmt.Scan(&num1)
					continue
				}
				result = lnVal(num1)
				graphFunc(num1, num2, operation)
			case "fact":
				fact := factVal(num1)
				fmt.Printf("\nResult: %s\n", fact.String())
				break
			case "abs":
				result = absVal(num1)

			case "trigo":
				fmt.Println("\nEnter trigonometric operation: (sin, cos, tan, cot, csc, sec) ")
				fmt.Scan(&trigonometry)
				switch trigonometry {
				case "sin":
					result = sinVal(num1)
					graphFunc(num1, num2, trigonometry)

				case "cos":
					result = cosVal(num1)
					graphFunc(num1, num2, trigonometry)

				case "tan":
					cosValue := cosVal(num1)
					if num1 == 90 || num1 == 270 || cosValue == 0 {
						fmt.Println("\nUndefined value.")
						fmt.Println("\nEnter the 1st Number: ")
						num1 = scanWithTypeCheck()
						continue
					} else {
						result = sinVal(num1) / cosValue
						graphFunc(num1, num2, trigonometry)
						break
					}

				case "cot":
					sinValue := sinVal(num1)
					if num1 == 0 || num1 == 180 || sinValue == 0 {
						fmt.Println("\nUndefined value.")
						fmt.Println("\nEnter the 1st Number: ")
						num1 = scanWithTypeCheck()
						continue
					} else {
						result = cosVal(num1) / sinValue
						graphFunc(num1, num2, trigonometry)
						break
					}

				case "csc":
					sinValue := sinVal(num1)
					if num1 == 0 || num1 == 180 || sinValue == 0 {
						fmt.Println("\nUndefined value.")
						fmt.Println("\nEnter the 1st Number: ")
						num1 = scanWithTypeCheck()
						continue
					} else {
						result = (1 / sinValue)
						graphFunc(num1, num2, trigonometry)
					}

				case "sec":
					cosValue := cosVal(num1)
					if num1 == 90 || num1 == 270 || cosValue == 0 {
						fmt.Println("\nUndefined value.")
						fmt.Println("\nEnter the 1st Number: ")
						num1 = scanWithTypeCheck()
						continue
					} else {
						result = (1 / cosVal(num1))
						graphFunc(num1, num2, trigonometry)
					}
				default:
					fmt.Println("\nInvalid trigonometric operation. Please enter a valid operation.")
					continue
				}
			default:
				fmt.Println("\nInvalid operation. Please enter a valid operation.")
				continue
			}

			if operation != "fact" {
				fmt.Printf("\nResult: %v\n", result)
			}
			break
		}
		if !isContinue() {
			break
		} else {
			clearResult(&result)
		}
	}
}
