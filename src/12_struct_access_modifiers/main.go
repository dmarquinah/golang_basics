package main

import (
	"fmt"
	pk "golang_basics/src/12_struct_access_modifiers/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	// The following lines will mark an error
	// The referenced struct is private
	/* var myOtherCar pk.carPrivate
	fmt.Println(myOtherCar) */

	pk.PrintMessage("Hello there!")

	// This line will also erroring out, because the method is private
	/* pk.printMessagePrivate("Hello there!") */
}
