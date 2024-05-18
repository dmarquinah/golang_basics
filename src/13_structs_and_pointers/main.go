package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

// We use pointer reference if we want to make changes in the struct itself
// If not, we would obtain a copy of the struct
func (myPc *pc) duplicateRAM() {
	myPc.ram *= 2
	fmt.Println("Duplicated RAM:", *myPc)
}

func main() {
	a := 50
	b := &a // Memory direction

	fmt.Println(a, b)
	fmt.Println(b, *b) // * points to the data held by the memory reference

	*b = 100 // All references to the memory space will change
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println("myPC:", myPC)

	myPC.ping()

	// myPC.ram = 20 --> This is redundant, use method with pointer instead
	myPC.duplicateRAM()
}
