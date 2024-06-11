package main

import (
	"fmt"
)

// wrapping native string type as customString type
type customString string

func (text customString) log() {
	// adding a custom method to the string type
	fmt.Println(text)
}

func main() {
	var name customString = "Maximum"
	name.log()
}
