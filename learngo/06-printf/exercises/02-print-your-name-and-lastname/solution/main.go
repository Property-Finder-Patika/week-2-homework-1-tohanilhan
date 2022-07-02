package main

import "fmt"

func main() {
	fmt.Printf("My name is %s and my lastname is %s.\n", "Tohan", "İlhan")

	// BONUS
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, "Tohan", "İlhan")
}
