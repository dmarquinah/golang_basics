package main

import "fmt"

func main() {
	// Showcasing mathematical operations
	const triangleBase = 5
	const triangleHeight = 16
	triangleArea := (triangleBase * triangleHeight) / 2
	fmt.Println("Triangle Area:", triangleArea)

	// Variable initialization
	x := 10
	y := 50

	// Sum expression
	result := x + y
	fmt.Println("Sum:", result)

	// Difference expression
	result = y - x
	fmt.Println("Difference:", result)

	// Multiplication expression
	result = x * y
	fmt.Println("Multiplication:", result)

	// Division expression
	result = y / x
	fmt.Println("Division:", result)

	// Module expression
	result = y % x
	fmt.Println("Module:", result)

	// Incremental
	x++
	fmt.Println("X incremented:", x)

	// Decremental
	y--
	fmt.Println("Y decremented:", y)
}
