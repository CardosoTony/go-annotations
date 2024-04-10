package main

import "fmt"

func main() {
	number := 2

	var pointer *int = nil
	fmt.Println(pointer)

	// Go does not have pointer arithmetic
	// pointer++

	pointer = &number // get pointer address
	fmt.Println(pointer)

	*pointer++ // dereferencing (fetching the value)
	fmt.Println(*pointer)

	number++

	fmt.Println(pointer, *pointer, number, &number)
}
