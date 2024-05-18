package main

import (
	"fmt"
	"sync"
)

func speak(text string, wg *sync.WaitGroup) {
	defer wg.Done() // the defer keyword will make the desired line of code to execute at the end
	fmt.Println(text)
}

// The main function is a routine itself
func main() {
	// WaitGroup is used to "save" and "liberate" desired executions by waiting on them
	var wg sync.WaitGroup // Alternative to handle concurrency, but hard to maintain

	// speak("Hello") // If this line is executed, the routine will end and the execution finishes, not showing the following line
	fmt.Println("Hello") // Fixed version from previous line that uses a simple console output

	wg.Add(1) // We're telling we have 1 remaining execution to wait

	// go speak("World") // Tries to use the concurrent thread used to execute the first line
	go speak("World", &wg) // Fixed version for previous line that uses a WaitGroup

	wg.Wait() // We want to wait the last concurrent execution to be finished

	wg.Add(1) // Added again if we want to provide a new concurrent execution

	// Anonymous function to demonstrate we still can provide new functionality, easier to check it by using the time.Sleep line
	go func(text string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(text)
	}("Good bye!", &wg)

	wg.Wait() // Waiting again for last execution

	//time.Sleep(time.Second * 1) // Not recommended, just a hack to show the concurrency without the WaitGroup
}
