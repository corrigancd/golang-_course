package main

import (
	"fmt"
	// "math"
)

func main() {
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	ebt, profit, ratio := calclutateValues(revenue, expenses, taxRate)
	printValues(ebt, profit, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Printf("Please enter value for %v: ", infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calclutateValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = (revenue - expenses)
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func printValues(ebt, profit, ratio float64) {
	fmt.Printf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
}
