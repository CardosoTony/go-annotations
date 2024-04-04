package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Integer numbers
	fmt.Println(1, 100, 1000)
	intNumber := 500
	fmt.Printf("Number %v is of '%T' type\n", intNumber, intNumber)
	fmt.Println("Number", intNumber, "is of type", reflect.TypeOf(intNumber))

	// Unsigned integer numbers (only positive numbers)
	// uint8, uint16, uint32, uint64

	var maxUintByte byte = math.MaxUint8
	fmt.Println("Max value for a byte is", maxUintByte)
	fmt.Println("Byte is of type:", reflect.TypeOf(maxUintByte))

	// Signed integer numbers (positive and negative numbers)
	// int8, int16, int32, int64

	minInt64 := math.MinInt64
	maxInt64 := math.MaxInt64
	fmt.Println("Min value for an int64 is", minInt64)
	fmt.Println("Max value for an int64 is", maxInt64)

	var unicodeReference rune = 'a'
	fmt.Println("Rune is of type", reflect.TypeOf(unicodeReference))
	fmt.Printf("The Unicode representation of 'a' is %v\n", unicodeReference)

	// Float numbers (float32 and float64)
	floatNumber := 23.99
	var floatNumber1 float32 = 51.99
	fmt.Printf("Number %v is of type '%T'\n", floatNumber, floatNumber)
	fmt.Println("Number", floatNumber1, "is of type", reflect.TypeOf(floatNumber1))

	// Boolean type
	boolean := true
	fmt.Printf("Boolean %v is of type '%T'\n", boolean, boolean)

	// String type
	str := "Hello World!"
	fmt.Printf("String %v is of type '%T'\n", str, str)
	fmt.Println("The length of the string is:", len(str))

	// Multiline string
	multilineString := `This is a multiline string.
  This is another line.`
	fmt.Println(multilineString)
	fmt.Println("The length of the multiline string is:", len(multilineString))

	// Character type (In Go, characters are represented as int numbers of Unicode)
	character := 'a'
	fmt.Printf("Character 'a' = %v is of type '%T' type\n", character, character)
	fmt.Println("Character", character, "is of type", reflect.TypeOf(character))
}
