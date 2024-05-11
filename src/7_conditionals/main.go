package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value1 := 1
	value2 := 3

	if value1 == 1 {
		fmt.Println("First value is 1")
	} else {
		fmt.Println("First value isn't 1")
	}

	// With and
	if value1 == 1 && value2 == 2 {
		fmt.Println("AND statement is true.")
	}

	// With or
	if value1 == 1 || value2 == 2 {
		fmt.Println("OR statement is true.")
	}

	// Handling errors
	value, err := strconv.Atoi("55")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	// Switch sentence
	// Used when want to use multiple conditionals

	switch myModule := 8 % 2; myModule {
	case 0:
		fmt.Println("Is even")
	default:
		fmt.Println("Is odd")
	}

	// Without condition
	value3 := -150
	switch {
	case value3 > 100:
		fmt.Println("Value3 greater than 100")
	case value3 < 0:
		fmt.Println("Value3 is negative")
	default:
		fmt.Println("No condition")
	}
}
