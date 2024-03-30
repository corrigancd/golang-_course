package main

import (
	"fmt"
	// "math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("How much did you earn: ")
	fmt.Scan(&revenue)
	fmt.Print("How much in expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("What is your % tax rate: ")
	fmt.Scan(&taxRate)

	ebt := (revenue - expenses)
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("EBT: €", ebt)
	fmt.Println("Profit: €", profit)
	fmt.Println("Ratio: ", ratio)

}
