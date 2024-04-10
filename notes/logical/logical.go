package main

import "fmt"

func purchases(job1, job2 bool) (bool, bool, bool) {
	buyItem1 := job1 && job2
	buyItem2 := job1 != job2 // 'exclusive OR' or 'XOR'
	buyItem3 := job1 || job2

	return buyItem1, buyItem2, buyItem3
}

func main() {
	buyItem1, buyItem2, buyItem3 := purchases(true, true)
	fmt.Printf("Item 1: %t, Item 2: %t, Item 3: %t, Item 3 negation: %t\n", buyItem1, buyItem2, buyItem3, !buyItem3)
}
