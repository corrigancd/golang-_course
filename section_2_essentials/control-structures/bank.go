package main

import "errors"
import "fmt"
import "os"
import "strconv"

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read account balance from file")
	}

	balanceText := string(data) // convert to float and return
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func getAmount(infoText string) float64 {
	var amount float64
	fmt.Printf("How much do you want to %v: ", infoText)
	fmt.Scan(&amount)
	return amount
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------------")
		panic("Can't continue, sorry")
	}

	fmt.Println("Wecome go Go Bank!")

	var exitBank bool = false

	// for { // this is an infinite loop, but you can use break to exit from it
	for !exitBank { // this is a boolean variable to control whether or not to exit from the loop
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)

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
