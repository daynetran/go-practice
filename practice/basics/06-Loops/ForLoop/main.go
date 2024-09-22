package main

import "fmt"

func main() {
	// basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// infinite for loop
	i := 0
	for {
		fmt.Println("Hello world")
		if i == 10 {
			break
		}
		i++
	}

	strings := []string{"Hello", "World", "!"}

	//
	for i, val := range strings {
		fmt.Printf("%d %s \n", i, val)
	}

	for i := range strings {
		fmt.Println(i)
	}

	for _, val := range strings {
		fmt.Println(val)
	}
}
