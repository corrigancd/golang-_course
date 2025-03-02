package main

import (
	"errors"
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // send a signal that the task is done
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // send a signal that the task is done
}

func greetWithError(phrase string, errorChan chan error) {
	fmt.Println("Hello!", phrase)
	errorChan <- errors.New("error cdtest errror")
}

func main() {
	dones := make([]chan bool, 5)
	errorChans := make([]chan error, 5)

	for i := range dones {
		dones[i] = make(chan bool)       // keep track of the channels so we can close them later on when all tasks are done
		errorChans[i] = make(chan error) // keep track of the channels so we can close them later on when all tasks are done
	}

	// start multiple goroutines to greet different phrases concurrently
	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])
	go greetWithError("I hope you're liking the course with Error!", errorChans[4])

	for index := range dones {
		select {
		case <-dones[index]:
			// fmt.Println("Done")
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// for _, done := range dones {
	// 	<-done
	// }

}
