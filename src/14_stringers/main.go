package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

// We override the function with name String to format the struct output
func (myPC pc) String() string {
	return fmt.Sprintf("I have %d GB of RAM, %d disk space and a %s brand PC\n", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	// Typical definition and print-out of struct
	myPc := pc{ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPc)
}
