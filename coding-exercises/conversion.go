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
	var intNum int
	// get the value from the user
	fmt.Println("For int variables: ")
	fmt.Println("Enter the number: ")
	fmt.Scanln(&intNum)
	// convert the value to float64
	fmt.Println("The float64 value of the intNumber is: ", float64(intNum))
	// convert the value to string
	fmt.Println("The string value of the intNumber is: ", strconv.Itoa(intNum))

	// create a variable for conversion
	var floatintNum float64
	// get the value from the user
	fmt.Println("For float variables: ")
	fmt.Println("Enter the floatNumber: ")
	fmt.Scanln(&floatintNum)
	// convert the value to int
	fmt.Println("The int value of the floatNumber is: ", int(floatintNum))
	// convert the value to string
	fmt.Println("The string value of the floatNumber is: ", strconv.FormatFloat(floatintNum, 'f', -1, 64))

	// create a variable for conversion
	var stringintNum string
	// get the value from the user
	fmt.Println("For string variables: ")
	fmt.Println("Enter the stringNumber: ")
	fmt.Scanln(&stringintNum)
	// convert the value to int
	intNum4, err := strconv.Atoi(stringintNum)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The int value of the stringNumber is: ", intNum4)
	// convert the value to float64
	intNum5, err := strconv.ParseFloat(stringintNum, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("The float64 value of the stringNumber is: ", intNum5)

	// create a variable for conversion
	var boolintNum bool
	// get the value from the user
	fmt.Println("For bool variables: ")
	fmt.Println("Enter the boolNumber: ")
	fmt.Scanln(&boolintNum)
	// convert the value to string
	fmt.Println("The string value of the boolNumber is: ", strconv.FormatBool(boolintNum))

}
