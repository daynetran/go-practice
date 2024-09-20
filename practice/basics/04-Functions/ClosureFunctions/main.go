package main

import "fmt"

func main() {
	l := 69
	w := 420

	func() {
		area := l * w
		fmt.Println(area)
	}()
}
