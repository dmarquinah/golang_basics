package main

import "fmt"

func main() {
	// Conditional For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------")

	// For while
	counter := 2
	for counter < 100 {
		counter *= 2
		fmt.Println(counter)
	}

	// Forever for
	// Not actually viable but can be defined like that
	/*
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/

	//Example with range
	numberList := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, value := range numberList {
		fmt.Printf("Position %d from list with value: %d \n", i, value)
	}
}
