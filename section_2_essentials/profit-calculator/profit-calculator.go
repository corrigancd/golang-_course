package main

import (
	"errors"
	"fmt"
	"os"
	// "math"
)

const financialsFile = "financials.txt"

func writeFinancialsToFile(ebt, profit, ratio float64) {
	// balanceText := fmt.Sprint(ebt, profit, ratio)
	balanceText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile(financialsFile, []byte(balanceText), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calclutateValues(revenue, expenses, taxRate)

	printValues(ebt, profit, ratio)
	writeFinancialsToFile(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Printf("Please enter value for %v: ", infoText)
	fmt.Scan(&userInput)

	if userInput < 0 {
		fmt.Println("You can't enter a negative value")
		return 0, errors.New("You can't enter a negative value")
	}

	return userInput, nil
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
