package main

import "fmt"

func main() {
	x := [5]int{0, 5, 10, 15, 20}

	// Standard for loop
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Getting key and value from range
	for k, v := range x {
		fmt.Println(k, " => ", v)
	}

	// Only getting value from range
	for _, v := range x {
		fmt.Println(v)
	}

	// Range and counter
	j := 0
	for range x {
		fmt.Println(x[j])
		j++
	}
}
