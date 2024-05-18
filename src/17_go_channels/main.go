package main

import "fmt"

// The channel argument can be specified to be as input or output: chan<- as sending only || <-chan as recieving only
func speak(text string, c chan<- string) {
	c <- text
	// text = <- c // if we want to recieve data from the channel
}

func main() {
	c := make(chan string, 1) // Length can be omitted so the number of channels is dynamic, but it can be costly

	fmt.Println("Hello!")

	go speak("Bye!", c)

	fmt.Println(<-c)

}
