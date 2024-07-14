package main

import (
	"fmt"
)

func main() {
	// a string array
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"

	// productNames[0] = "Laptop"
	// productNames[1] = "Phone"
	// productNames[3] = "Mouse"

	// a number array
	prices := []float64{10.5, 15.7, 8.2, 12.9}

	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[2])

	// Slices are dynamic arrays that can grow and shrink

	// featuredPrices := prices[:3] // start at the beginning to index 3
	// featuredPrices := prices[1:] // continue from index 1
	// featuredPrices := prices[1:4] // the maximum index is the maximum index plus one
	featuredPrices := prices[1:3] // starting and ending (the start is included, the end is excluded)
	fmt.Println(featuredPrices)
	fmt.Printf("The second and third items from the prices list in an array: %v\n", featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Printf("Just the first item from the feature prices list in an array: %v\n", highlightedPrices)

	featuredPrices[0] = 10000
	fmt.Printf("Mutating an item in the slice, will also change the original array: %v\n", prices)

	fmt.Printf("Printing the length of array: %v\n", len(featuredPrices)) //length of an array, this is the length of actual items in the slice
	fmt.Printf("Printing the length of array: %v\n", cap(featuredPrices)) //capacity of an array, this is a way to determine the length of an array/slice, but we don't have to have all of the elements defined, some can be empty

	discountPrices := []float64{101.99, 100.99, 99.99}

	// the spread syntax is used to extract the individual values from the discount prices array
	allPrices := append(prices, discountPrices...)
	fmt.Println(allPrices)
}
