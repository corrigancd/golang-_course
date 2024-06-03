package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func getAmount(infoText string) float64 {
	var amount float64
	fmt.Printf("How much do you want to %v: ", infoText)
	fmt.Scan(&amount)
	return amount
}

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------------")
		panic("Can't continue, sorry")
	}

	fmt.Println("Wecome go Go Bank!")
	fmt.Println("Reach us 24/7 on:", randomdata.PhoneNumber())

	var exitBank bool = false

	// for { // this is an infinite loop, but you can use break to exit from it
	for !exitBank { // this is a boolean variable to control whether or not to exit from the loop
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("Your balance is", accountBalance)

		case 2:
			amount := getAmount("deposit")
			if amount <= 0 {
				fmt.Println("You can't deposit a negative amount")
				continue
			}
			accountBalance += amount
			fmt.Println("Your balance is", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			amount := getAmount("withdraw")
			if amount >= accountBalance {
				fmt.Println("You can't withdraw more than you have in your account")
				continue
			}
			accountBalance -= amount
			fmt.Println("Your balance is", accountBalance)

		default:
			fmt.Println("Goodbye")
			exitBank = true
		}
	}
	fmt.Println("Thanks for choosing our bank")
}
