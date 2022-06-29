// Example usage of iota
package main

import "fmt"

// create 4 variables with iota
const (
	a int = iota
	b
	c
	d
)

// create 4 variables with iota starting from 10
const (
	e int = iota + 10
	f
	g
	h
)

// create 4 variables with iota and increment it by 3
const (
	i int = iota * 3
	j
	k
	l
)

// create 4 variables with iota and decrement it by 3
const (
	m int = iota * -3
	n
	o
	p
)

func main() {
	// print the values of the variables
	fmt.Println("Default iota values starts from 0 and increments by 1: ")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// print the values of the variables
	fmt.Println("But you can manually set the starting value of iota by using the '+' operator: ")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	// print the values of the variables
	fmt.Println("You can also change the increment value of iota by using the '*' operator: ")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	// print the values of the variables
	fmt.Println("You can also change the decrement value of iota by using the '*-' operator: ")
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)

}
