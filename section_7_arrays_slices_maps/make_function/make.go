package main

import (
	"fmt"
)

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// make creates and array with 2 empty elements and a capacity of 5 (setting the capacity is more efficient than using the dynamic array)
	userNames := make([]string, 2, 5)
	// here we append 3 elements to the dynamic array
	userNames = append(userNames, "John Doe", "Jane Smith", "Bob Johnson")
	// so when we print, we get the two empty elements, then the three appended elements
	fmt.Println(userNames)

	// Now we can assign to the first two elements
	userNames[0] = "test name 1"
	userNames[1] = "test name 2"
	fmt.Println(userNames)

	// We can also use the make function when creating a map, for efficiency
	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.8 // go will have to allocate more memory to the map behind the scenes for this

	floatMap.output(courseRatings)

	// Iterating over arrays
	for index, value := range userNames {
		fmt.Println("index: ", index, "value: ", value)
	}

	// Iterating over maps
	for key, value := range courseRatings {
		fmt.Println("key: ", key, "value: ", value)
	}

}
