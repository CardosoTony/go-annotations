package main

import "fmt"

func main() {
	var a int     // 0
	var b float64 // 0.0
	var c bool    // false
	var d string  // ""
	var e *int    // nil

	fmt.Printf("a:%v\nb:%v\nc:%v\nd:%q\ne:%v\n", a, b, c, d, e)
}
