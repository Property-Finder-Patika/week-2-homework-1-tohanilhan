// Example of scope
package main

import "fmt"

// Global variable declaration
var x int = 5

func main() {

	fmt.Println("value of x as a global variable = ", x)

	// Local variable declaration in main function
	var x int = 7
	var y int = 12
	var z int

	fmt.Printf("Value of x after manipulating its value in main() function = %d\n", x)
	z = sum(x, y)
	fmt.Printf("Value of z in main() function = %d\n", z)
}

// Function to add two integers
func sum(a, b int) int {
	fmt.Printf("Value of a(x) in sum() function = %d\n", a)
	fmt.Printf("Value of b(y) in sum() function = %d\n", b)

	return a + b
}
