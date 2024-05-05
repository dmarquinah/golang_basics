package main

import "fmt"

func regularFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnedFunction(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * a
}

func main() {
	regularFunction("Executing regular function!")
	tripleArgument(1, 2, "triple")

	value := returnedFunction(5)
	fmt.Println("Returned value:", value)

	value1, value2 := doubleReturn(9)
	fmt.Println("Double returned values:", value1, value2)

	// If not interested in second argument just use underscore (_)
	_, value3 := doubleReturn(7)
	fmt.Println("One of double returned values:", value3)
}
