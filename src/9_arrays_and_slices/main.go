package main

import "fmt"

func main() {
	// Defining Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Slice methods

	// In slice[a:b]
	// a is inclusive (going to be included); b is exclusive (going to be excluded)
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
