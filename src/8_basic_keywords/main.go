package main

import "fmt"

func main() {
	// Defer
	// Useful to close connections, since it's executed at the end of the regular code
	defer fmt.Println("Hello")
	fmt.Println("World")

	// Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("It's 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Breaking in 8")
			break
		}
	}
}
