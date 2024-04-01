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

	// format and strore string as value
	formattedFv := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRfv := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futreRealValue)
	fmt.Print(formattedFv, formattedRfv) // log already formatted string to console

	// format and print directly to terminal
	fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futreRealValue)

}
