package main

import "fmt"

func main() {
	// Constants declaration
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Integer variables declaration
	base := 12
	var height int = 14
	var area int

	fmt.Println("base:", base)
	fmt.Println("height:", height)
	fmt.Println("area:", area)

	// Zero values
	// Default values defined by type
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Calculating square are
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area is:", squareArea)
}
