package main

import (
	"fmt"
	"math"
	"math/big"
)

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
func sinVal(num float64) float64 {
	return math.Sin(degreesToRadians(num))
}
func cosVal(num float64) float64 {
	return math.Cos(degreesToRadians(num))
}
func degreesToRadians(degree float64) float64 {
	return degree * math.Pi / 180
}
