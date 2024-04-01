package main

import "fmt"

func main() {
	// 'Print' displays on the same line
	fmt.Print("Hello ")
	fmt.Print("World!\n")

	// 'Println' displays on a new line
	fmt.Println("Hello ")
	fmt.Println("World!")

	x := 3.14159
	// fmt.Println("The value of x is:" + x) // Error message, cannot concatenate string with numbers

	// solution 1
	xString := fmt.Sprint(x)
	fmt.Println("The value of x is: " + xString)

	// solution 2
	fmt.Println("The value of x is:", x)

	// solution 3
	fmt.Printf("The value of x is: %f\n", x)
	fmt.Printf("The value of x is: %.2f\n", x)

	intNumber := 4
	floatNumber := 1.999
	boolean := false
	stringType := "hello"
	fmt.Printf("\nint = %d,\nfloat = %f,\nfloat = %.1f,\nboolean = %t,\nstring = %s\n", intNumber, floatNumber, floatNumber, boolean, stringType)
	fmt.Printf("\nint: %v,\nfloat: %v,\nboolean: %v,\nstring: %v\n", intNumber, floatNumber, boolean, stringType)
}
