// Example of conversions among basic types
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// create conversion examples for the following basic types:
	// int, float64, string, bool
	// create a variable for conversion
	var num int
	// get the value from the user
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	// convert the value to float64
	fmt.Println("The float64 value of the number is: ", float64(num))
	// convert the value to string
	fmt.Println("The string value of the number is: ", strconv.Itoa(num))

	// create a variable for conversion
	var num2 float64
	// get the value from the user
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num2)
	// convert the value to int
	fmt.Println("The int value of the number is: ", int(num2))
	// convert the value to string
	fmt.Println("The string value of the number is: ", strconv.FormatFloat(num2, 'f', -1, 64))

	// create a variable for conversion
	var num3 string
	// get the value from the user
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num3)
	// convert the value to int
	num4, err := strconv.Atoi(num3)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The int value of the number is: ", num4)
	// convert the value to float64
	num5, err := strconv.ParseFloat(num3, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The float64 value of the number is: ", num5)

	// create a variable for conversion
	var num6 bool
	// get the value from the user
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num6)
	// convert the value to string
	fmt.Println("The string value of the number is: ", strconv.FormatBool(num6))

}
