package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

// Use lowercase for private struct and fields
type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}

// Methods can be private as well using lowercase
func printMessagePrivate(text string) {
	fmt.Println(text)
}
