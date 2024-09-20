package main

import "fmt"

func main() {
	x := true
	if x {
		fmt.Println("True")
	}

	y := 100
	if y > 80 {
		fmt.Println("Greater than 80")
	} else {
		fmt.Println("Less than or equal to 80")
	}

	grade := 85
	if grade > 90 {
		fmt.Println("Grade is A")
	} else if grade >= 70 && grade < 90 {
		fmt.Println("Grade is not A, but passing")
	} else {
		fmt.Println("Grade is failing.")
	}

	if a := 10; a > 5 {
		fmt.Println("Greater than 5")
	}
}
