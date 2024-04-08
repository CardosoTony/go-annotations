package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 5.4
	y := 3
	fmt.Println(x / float64(y))

	floatNumber := 7.9
	intNumber := int(floatNumber)
	fmt.Println(intNumber)

	// Convert int to string
	// Warning
	fmt.Println("Wrong way to convert: " + string(97))

	// Correct way
	fmt.Println("Correct way to convert: " + strconv.Itoa(97))

	// Convert string to int
	num, _ := strconv.Atoi("97")
	result := num - 7
	fmt.Printf("%v - 7 = %v\n", num, result)
}
