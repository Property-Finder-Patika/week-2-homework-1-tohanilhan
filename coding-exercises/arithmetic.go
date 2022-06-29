// Basic calculator
package main

import "fmt"

func main() {
	// create 2 variables for storing the numbers
	var num1, num2 int
	// create a variable for choice the result
	var choice int
	// get the values from the user
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)
	// select the operation to be performed
	fmt.Println("Select the operation to be performed: ")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	// get the choice from the user
	fmt.Scanln(&choice)
	// print the result based on the choice
	switch choice {
	case 1:
		fmt.Println("The sum of the numbers is: ", num1+num2)
	case 2:
		fmt.Println("The difference of the numbers is: ", num1-num2)
	case 3:
		fmt.Println("The product of the numbers is: ", num1*num2)
	case 4:
		fmt.Println("The quotient of the numbers is: ", num1/num2)
	default:
		fmt.Println("Invalid choice")
	}

}
