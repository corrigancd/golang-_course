package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	// var investmentAmount, years float64 = 1000, 10 // because we explicitly set it to float64, all variables here are floats
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futreRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futreRealValue)
}
