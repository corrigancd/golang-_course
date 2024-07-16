package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 10, 15, 20, 25}
	doubled := doubleNumbers(&numbers)

	fmt.Println(doubled)

	for _, val := range numbers {
		fmt.Println(double(val))
	}

	// returning a function from a function and passing it to another function
	fmt.Println(transformNumbers(&numbers, getTransformerFunction(&numbers)))
	fmt.Println(transformNumbers(&moreNumbers, getTransformerFunction(&moreNumbers)))

	// passing a function as a parameter to a function
	fmt.Println(transformNumbers(&numbers, triple))

	// anonymous function example
	anonymousFunctionTransform := transformNumbers(&numbers, func(number int) int {
		return number * 4
	})
	fmt.Println(anonymousFunctionTransform)

	// closure example
	closure := createTransformerFunction(5)
	fmt.Println(transformNumbers(&numbers, closure))

}

func doubleNumbers(numbers *[]int) []int {
	dnumbers := []int{}
	for _, val := range *numbers {
		dnumbers = append(dnumbers, val*2)
	}
	return dnumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// transformNumbers takes a slice of integers and a function as parameters,
// applies the function to each element in the slice, and returns a new slice with the transformed values.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformedValues := []int{}
	for _, val := range *numbers {
		transformedValues = append(transformedValues, transform(val))
	}
	return transformedValues
}

// returning the double function from this function
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func createTransformerFunction(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
