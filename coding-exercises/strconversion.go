// Conversion example using strconv
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// create a variable for conversion
	var stringNum string
	// get the value from the user
	fmt.Println("Enter the stringNumber: ")
	fmt.Scanln(&stringNum)
	// convert the value to int
	intNum4, err := strconv.Atoi(stringNum)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The int value of the stringNumber is: ", intNum4)
	// convert the value to float64
	intNum5, err := strconv.ParseFloat(stringNum, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The float64 value of the stringNumber is: ", intNum5)

}
