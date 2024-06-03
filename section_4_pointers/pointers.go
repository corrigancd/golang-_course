package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	age := 32 // regular variable

	var agePointer *int // verbose to show how pointer variable is declared
	agePointer = &age   // pointer to the age variable

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", agePointer)
	fmt.Println("Age Pointer de-referenced: ", *agePointer)
	fmt.Println("Adult years: ", getAdultYears(agePointer))

	fmt.Println("Proving that the original value of age has not yet been overwritten: ", age)

	//assgining in function and logging to console
	assignAdultYearsFromTwentyOneToPointer(agePointer)
	fmt.Println("Adult years: ", *agePointer)
}

// An example of how to pass a pointer to a function
// Note you should not do this, because the integer value
// is tiny. But for larger data types, this may make sense
// example of assinging value to a point and returning that
func getAdultYears(age *int) int {
	if *age >= 18 {
		return *age - 18
	}
	return 0

}

func assignAdultYearsFromTwentyOneToPointer(age *int) {
	// example of assigning the result of a some logic back to a pointer
	// This is usefult for mutating data inside of a function (if you really want to do that)

	if *age >= 18 {
		*age = *age - 21
	}
}
