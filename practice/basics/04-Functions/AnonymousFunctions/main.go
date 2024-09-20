package main

import "fmt"

// Package-level anonymous function
var (
	area = func(l int, w int) int {
		return l * w
	}
)

func main() {
	a := area(6, 9)
	fmt.Println(a)

	// Function-level anonymous function
	volume := func(l int, w int, h int) int {
		return l * w * h
	}
	v := volume(6, 9, 420)
	fmt.Println(v)
}
