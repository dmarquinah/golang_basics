package main

import "fmt"

type shape2D interface {
	getArea() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	return s.base * s.base
}

func (r rectangle) getArea() float64 {
	return r.base * r.height
}

func calculate(f shape2D) {
	fmt.Println("Area:", f.getArea())
}

func main() {
	mySquare := square{base: 5}
	myRectangle := rectangle{base: 2, height: 7}

	// If we don't use interfaces we would have to do:
	// mySquare.getArea()
	// myRectangle.getArea()

	// Using the interface
	calculate(mySquare)
	calculate(myRectangle)

	// We can also have an interface list to define different type of data
	myInterface := []interface{}{"Hello", 12, 4.95}
	fmt.Println(myInterface...)
}
