package main

import (
	"flag"
	"fmt"
	_ "math"
	"strconv"
)

// Operation functions.
func addVal(num1, num2 float64) float64 {
	return num1 + num2
}
func subVal(num1, num2 float64) float64 {
	return num1 - num2
}
func mulVal(num1, num2 float64) float64 {
	return num1 * num2
}
func divVal(num1, num2 float64) float64 {
	return num1 / num2
}

// The function for User's request to continue or exit the program
func isContinue() bool {

	var str string
	for {
		fmt.Println("\nDo you want to continue or exit the program? (c/e): ")
		fmt.Scan(&str)
		if str == "c" {
			return true
		} else if str == "e" {
			return false
		} else {
			fmt.Println("\nJust press c/n please.")
		}
	}
}

// Type check with scan for user input.
func scanWithTypeCheck() float64 {

	var num string
	for {

		fmt.Scan(&num)
		var num, err = strconv.ParseFloat(num, 64)
		if err == nil {
			return num
		} else {
			//strconv.ParseFloat: parsing "_": invalid syntax
			fmt.Println("\nInvalid input. Please enter a valid number.\n")

		}

	}
}

// The function that allows the user to clear the result and start the operation from scratch.
func clearResult(result *float64) {

	var clr string

	fmt.Println("\nIf you want a clear result, type 'clear', otherwise press any key.")
	fmt.Scan(&clr)

	if clr == "clear" {
		fmt.Println("\nClearing the result.")
		*result = 0
	} else {
		fmt.Println("\nYou choose to continue with result.")
	}
}

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
				fmt.Printf("\nResult: %g \n", result)

			case "-":
				result = subVal(num1, num2)
				fmt.Printf("\nResult: %g \n", result)

			case "*":
				result = mulVal(num1, num2)
				fmt.Printf("\nResult: %g \n", result)

			case "/":
				if num2 == 0 {
					fmt.Println("\nCan't division by zero.")
					continue
				}
				result = divVal(num1, num2)
				fmt.Printf("\nResult: %g \n", result)

			default:

				fmt.Println("\nInvalid operation. Please enter a valid operation.\n")

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

func main() {

	type1 := flag.Bool("type1", false, "Basic Calculator")
	type2 := flag.Bool("type2", false, "Scientific Calculator")

	flag.Parse()

	if *type1 {
		bscCalc()

	} else if *type2 {

		fmt.Println("Scientific Calculator not yet implemented.")

	} else {
		fmt.Println("Invalid operation. Use -type1 for basic or -type2 for scientific calculator.")
	}

}
