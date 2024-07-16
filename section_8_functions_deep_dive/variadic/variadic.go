package main

import (
	"fmt"
)

func main() {
	// regular pass array to a function
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(&numbers)
	fmt.Println("Sum of numbers:", sum)

	// pass individual variables to a function
	startingValue := 10
	sumVariadic := sumupVariadic(startingValue, 1, 2, 3, 4, 5)
	fmt.Println("Variadic sum:", sumVariadic)

	// pass an array as variables to a function
	sumVariadicAsArray := sumupVariadic(startingValue, numbers...)
	fmt.Println("Variadic sum where numbers are spread:", sumVariadicAsArray)

}

func sumup(numbers *[]int) int {
	sum := 0
	for _, num := range *numbers {
		sum += num
	}
	return sum
}

func sumupVariadic(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, num := range numbers {
		sum += num
	}
	return sum
}
