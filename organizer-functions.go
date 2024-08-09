package main

import (
	"fmt"
	"strconv"
)

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
			fmt.Println("\nJust press c/e please.")
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
			fmt.Println("\nInvalid input. Please enter a valid number:")

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
