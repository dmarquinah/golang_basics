package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"Hello", "there", "my", "friend"}

	for i, value := range slice {
		fmt.Println(i, value)
	}

	// Also can be
	/*
		for _, value := range slice {
			fmt.Println(value)
		}

		for i := range slice {
			fmt.Println(i)
		}
	*/
	checkPalindrome("rats live on no evil star")
}

func checkPalindrome(text string) {
	// Also counting spaces to check
	slicedText := strings.Split(text, "")
	size := len(slicedText)
	isPalindrome := false

	// Another approach - Reverse the text
	for i, value := range slicedText {
		if value != slicedText[size-i-1] {
			break
		}

		if i > size/2 {
			isPalindrome = true
			break
		}
	}

	if isPalindrome {
		fmt.Println(text, "is palindrome")
	} else {
		fmt.Println(text, "is not palindrome")
	}

}
