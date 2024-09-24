package main

import "fmt"

func main() {
	var intArr [5]int

	// Assign values
	intArr[0] = 10
	intArr[1] = 2

	// Access values
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])

	// Initialize array with array literals
	x := [5]int{0, 5, 10, 15, 20}
	var y [5]int = [5]int{0, 5, 10, 15, 20}

	fmt.Println(x)
	fmt.Println(y)

	// Initialize array with ellipses
	k := [...]int{10, 20, 30}
	fmt.Println(len(k))

	// Initialize values for specific elements
	a := [5]int{1: 1, 3: 25}
	fmt.Println(a)
}
