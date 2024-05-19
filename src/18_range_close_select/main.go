package main

import "fmt"

func saveMessage(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "First message"
	c <- "Second message"

	// len(c) for the current number of used channels
	// cap(c) for the total capacity of channels that can be held
	fmt.Println(len(c), cap(c))

	// Range and Close

	// For closing a channel when it's not going to be used anymore because it's full for example
	close(c) // but also doesn't matter if it still have space to use other channel
	// c <- "Third message" // If we want to add a new one when no space, program will panic

	// Range to show all content in channel
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go saveMessage("Message 1", email1)
	go saveMessage("Message 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("message from 1st email:", m1)
		case m2 := <-email2:
			fmt.Println("message from 2nd email:", m2)
		}
	}
}
