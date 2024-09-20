package main

import "fmt"

func main() {
	// APPROACH ONE:
	// 1. Declare variables
	var i int
	var s string

	// 2. Initialize variables
	i = 32
	s = "Hello world"

	fmt.Println(i)
	fmt.Println(s)

	// APPROACH TWO:
	// Declare and initialize variable in one line
	var k int = 35
	fmt.Println(k)

	// APPROACH THREE:
	// Short variable declaration
	j := 50
	fmt.Println(j)

	// Multiple variable declaration
	firstName, lastName := "Dayne", "Tran"
	fmt.Println(firstName + " " + lastName)

	// Variable declaration block
	var (
		name     = "Donald Duck"
		age  int = 50
	)
	fmt.Println(name)
	fmt.Println(age)
}
