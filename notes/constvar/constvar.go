package main

import (
	"fmt"
	m "math" // allow renaming the imported package
)

func main() {
	const PI float64 = 3.141592653589 // explicitly typed
	var radius = 3.2                  // automatic type inference, same as Rust

	// another way to create a variable and assign it value
	// short declaration
	area := PI * m.Pow(radius, 2)
	fmt.Printf("The area of the circle with radius %f is %.2f\n", radius, area)

	// declare multiple constants and variables in a single declaration
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// declare multiple constants and variables in a single line
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Hello"
	fmt.Println(g, h, i)
}
