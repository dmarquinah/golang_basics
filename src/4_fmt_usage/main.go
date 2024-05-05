package main

import "fmt"

func main() {
	// Declaring variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Printing with new line: Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printing with format: Printf
	name := "Diego"
	experience := 3

	fmt.Printf("%s has more than %d years of experience\n", name, experience)

	// In case we don't know the type, we can use a wildcard (%v) for any operant
	fmt.Printf("%v has more than %v years of experience\n", name, experience)

	// Returning a string: SPrintf

	message := fmt.Sprintf("%s has more than %d years of experience", name, experience)
	fmt.Println("Message from string:", message)

	// Obtaining data types
	fmt.Printf("helloMessage Type: %T\n", helloMessage)
	fmt.Printf("experience type: %T\n", experience)

}
