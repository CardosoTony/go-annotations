package main

import "fmt"

func main() {
	x := 3
	y := 5

	// only postfix
	x++ // relative 'x += 1' or 'x = x + 1'
	fmt.Println(x)

	y-- // relative 'y -= 1' or 'y = y - 1'
	fmt.Println(y)
}
