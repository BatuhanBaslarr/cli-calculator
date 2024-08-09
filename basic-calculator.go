package main

import "fmt"

func bscCalc() {
	var operation string
	var num1, num2, result float64

	for {

		if result == 0 {
			fmt.Println("\nEnter 1st Number: ")
			num1 = scanWithTypeCheck()
		} else {
			num1 = result
		}

		for {
			fmt.Println("\nChoose Your Operation(+,-,*,/):  ")
			fmt.Scan(&operation)

			switch operation {

			case "+", "-", "*", "/":
				fmt.Println("\nEnter 2nd Number: ")
				num2 = scanWithTypeCheck()
			default:
				fmt.Println("\nInvalid operation. Please enter a valid operation.\n")
				continue
			}

			switch operation {

			case "+":
				result = addVal(num1, num2)
			case "-":
				result = subVal(num1, num2)
			case "*":
				result = mulVal(num1, num2)
			case "/":
				if num2 == 0 {
					fmt.Println("\nCan't division by zero.")
					continue
				}
				result = divVal(num1, num2)

			default:
				fmt.Println("\nInvalid operation. Please enter a valid operation.\n")
			}

			fmt.Printf("\nResult: %g\n", result)

			break
		}
		if !isContinue() {

			break
		} else {
			clearResult(&result)

		}
	}
}
