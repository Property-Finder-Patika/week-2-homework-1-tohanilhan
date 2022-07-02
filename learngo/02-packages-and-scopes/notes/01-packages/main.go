package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// You can access functions from other files which are in the same package

	// Here, `main()` can access `bye()` and `hey()`

	// It's because bye.go, hey.go and main.go are in the main package.

	// If you open the whole week-2-homework folder in the same editor,
	// you have to open the 01-packages file in the seperate editor to eliminate the error

	bye()
	hey()
}
