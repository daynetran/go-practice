package main

import "fmt"

func main() {
	i := 10
	k := 20
	z := 30

	// Arithmetic Operators
	fmt.Printf("i + k = %d\n", i+k)
	fmt.Printf("i - k = %d\n", i-k)
	fmt.Printf("i * k = %d\n", i*k)
	fmt.Printf("i / k = %d\n", i/k)
	fmt.Printf("i mod k = %d\n", i%k)

	// Comparison Operators
	fmt.Println(i == k)
	fmt.Println(i != k)
	fmt.Println(i < k)
	fmt.Println(i <= k)
	fmt.Println(i > k)
	fmt.Println(i >= k)

	// Logical Operators
	fmt.Println(i < z && i > k)
	fmt.Println(i < z || i > k)
	fmt.Println(!(i == z && i > k))

	// Assigment Operators
	var x, y = 15, 25
	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)
}
