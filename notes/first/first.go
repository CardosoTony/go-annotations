// Executable programs start with the 'main package'
package main

/*
Go code is organized into packages,
and to use them, it's necessary to declare one or more imports.
*/
import "fmt" // when it's just one import

// when there are multiple imports
/*
import (
	"fmt"
	"math"
)
*/

// Main function of Go
func main() {
	fmt.Println("Hello World!")
}
