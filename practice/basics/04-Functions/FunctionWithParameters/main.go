package main

import "fmt"

func main() {
	hello("Dayne")
	add(20, 30)
}

func hello(s string) {
	fmt.Printf("Hello, %s\n", s)
}

func add(x int, y int) {
	fmt.Println(x + y)
}
