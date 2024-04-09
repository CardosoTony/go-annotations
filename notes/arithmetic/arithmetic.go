package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5
	b := 3

	fmt.Println("Sum => ", a+b)
	fmt.Println("Difference => ", a-b)
	fmt.Println("Product => ", a*b)
	fmt.Println("Quotient => ", a/b)
	fmt.Println("Modulus => ", a%b)

	fmt.Println()

	// bitwise
	fmt.Println("Bitwise AND => ", a&b)
	fmt.Println("Bitwise OR => ", a|b)
	fmt.Println("Bitwise XOR => ", a^b)
	fmt.Println("Bitwise NOT => ", ^a)

	fmt.Println()

	// Operations using math
	fmt.Println("Min value between two numbers =>", math.Min(float64(a), float64(b)))

	c := 5.0
	d := 2.0
	fmt.Println("Max value between two numbers =>", math.Max(c, d))
	fmt.Println("Min value between two numbers =>", math.Min(c, d))
	fmt.Println("Power =>", math.Pow(c, d))
}
