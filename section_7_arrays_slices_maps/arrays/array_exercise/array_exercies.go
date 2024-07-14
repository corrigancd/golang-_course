package main

import (
	"fmt"
)

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func main() {
	// 1
	hobbies := [3]string{"football", "fishing", "diving"}
	fmt.Println(1, hobbies)

	// 2
	fmt.Println(2, hobbies[0])
	fmt.Println(2, hobbies[1:3])

	// 3
	someHobbiesMethod1 := hobbies[0:2]
	fmt.Println(3, someHobbiesMethod1)
	someHobbiesMethod2 := hobbies[:2]
	fmt.Println(3, someHobbiesMethod2)

	// 4
	fmt.Println(4, cap(someHobbiesMethod1))
	someHobbiesLast := someHobbiesMethod1[1:3]
	fmt.Println(4, someHobbiesLast)

	// 5
	goals := []string{"learn", "teach"}
	fmt.Println(5, goals)

	// 6
	goals[1] = "really teach"
	goals = append(goals, "understand")
	fmt.Println(6, goals)

	// 7
	type Product struct {
		title string
		id    int
		price float64
	}

	products := []Product{{title: "Product", id: 0, price: 10.99}, {title: "Product2", id: 1, price: 11.99}}
	products = append(products, Product{title: "Product3", id: 2, price: 12.99})
	fmt.Println(7, products)
}
