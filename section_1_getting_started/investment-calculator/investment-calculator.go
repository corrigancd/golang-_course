package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	years := 10.0
	expectedReturnRate := 5.5

	// var investmentAmount, years float64 = 1000, 10 // because we explicitly set it to float64, all variables here are floats
	var investmentAmount float64
	// & is a pointer to the variable, more later in the course!
	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// the float64 type is inferred by the compiler
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futreRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print("Future Value: %v\n", futureValue)
	fmt.Println("Future Value (adjusted for inflation):", futreRealValue)

}
