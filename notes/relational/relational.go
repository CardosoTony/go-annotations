package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String", "example" == "example")
	fmt.Println("!=", 5 != 3)
	fmt.Println("<", 5 < 3)
	fmt.Println(">", 5 > 3)
	fmt.Println("<=", 5 <= 3)
	fmt.Println(">=", 5 >= 3)
	fmt.Println("==", 5 == 3)

	date1 := time.Unix(0, 0)
	date2 := time.Unix(0, 0)

	fmt.Println("Date:", date1 == date2)
	fmt.Println("Date:", date1.Equal(date2))

	type Person struct {
		Name string
	}

	p1 := Person{Name: "John"}
	p2 := Person{Name: "John"}

	fmt.Println("Person:", p1 == p2)
}
