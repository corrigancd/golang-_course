package main

import (
	"fmt"
)

func main() {
	const factor int = 6
	fmt.Println(factorialIterative(factor))

	result := factorialRecursive(factor)
	fmt.Println(result)
}

func factorialIterative(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func factorialRecursive(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorialRecursive(number-1)
}
