package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
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
func sqrtVal(num float64) float64 {
	if num < 0 {

		fmt.Println("Can not take the square root of a negative numbers.")
		return 0
	} else {
		return math.Sqrt(num)
	}
}
func powVal(num1, num2 float64) float64 {
	return math.Pow(num1, num2)

}
func sqrVal(num float64) float64 {
	return num * num
}
func logVal(num1, num2 float64) float64 {
	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Logarithm base and argument must be greater than zero")
		return 0
	} else {
		return math.Log(num1) / math.Log(num2)
	}
}
func lnVal(num float64) float64 {
	if num <= 0 {
		fmt.Printf("Ln logarithm argument must be greater than zero")
		return 0
	} else {
		return math.Log(num)
	}
}
func factVal(num float64) *big.Int {
	if num < 0 {
		fmt.Println("Factorial of a negative number is not defined")
		return big.NewInt(0)
	}

	fact := big.NewInt(1)
	n := big.NewInt(int64(num))

	for i := big.NewInt(1); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		fact.Mul(fact, i)
	}

	return fact
}

func absVal(num float64) float64 {
	return math.Abs(num)
}
func sinVal(num float64) float64 {
	return math.Sin(num)
}
func cosVal(num float64) float64 {
	return math.Cos(num)
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

// Scientific Calculator
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
			fmt.Print("\nEnter operation (+, -, *, /, sqrt, pow, sqr, log, ln, fact, abs, trigo): ")
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
				result = logVal(num1, num2)
			case "ln":
				result = lnVal(num1)
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
				case "cos":
					result = cosVal(num1)
				case "tan":
					result = sinVal(num1) / cosVal(num1)
				case "cot":
					result = cosVal(num1) / sinVal(num1)
				case "csc":
					result = 1 / sinVal(num1)
				case "sec":
					result = 1 / cosVal(num1)
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

// Basic Calculator
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

// Main Function
func main() {

	type1 := flag.Bool("type1", false, "Basic Calculator")
	type2 := flag.Bool("type2", false, "Scientific Calculator")

	flag.Parse()

	if *type1 {
		bscCalc()

	} else if *type2 {

		sciCalc()

	} else {
		fmt.Println("Invalid operation. Use -type1 for basic or -type2 for scientific calculator.")
	}

}
