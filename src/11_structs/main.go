package main

import "fmt"

type Car struct {
	brand string
	year  int
}

func main() {
	myCar := Car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Second type of intantiation
	var otherCar Car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
