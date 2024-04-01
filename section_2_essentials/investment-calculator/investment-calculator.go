package main

import (
	"fmt"
	"math"
)

// variables can be declared at the top level of files
const inflationRate = 6.5

func main() {
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

	// You can print with multiple lines using backticks
	fmt.Printf(`Future Value: %.1f
	Future Value (adjusted for inflation): %.1f
	
	`, futureValue, futreRealValue)

	// example of calling the declared function below
	outputText(formattedFv)

	// futureValue and futureRealValue returned from the function (you separate them with a comma)
	funcFv, funcFrv := calculateFutureValues(investmentAmount, expectedReturnRate, years)
}

// Example of how to declare any function
func outputText(text string) {
	fmt.Println("Custom function: ")
	fmt.Println(text)
}

// function returns two values: (futureValue, futureRealValue) (float64, float64)
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) { // specifying one float64 type at the end means that all of the variables before this will also be float64
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
