package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getBalanceFromFile() (float64, error) {

	data, err := os.ReadFile("controled-structures/balance.txt")

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	convertedNum, convertError := strconv.ParseFloat(balanceText, 64)

	if convertError != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return convertedNum, nil

}

func writeBalanceToFile(balance float64) {

	balanceText := fmt.Sprint(balance)
	os.WriteFile("controled-structures/balance.txt", []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to GO Bank!")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("Your balance is", accountBalance)
			fmt.Println("------------------------")

		case 2:
			{

				fmt.Print("Your Deposit:")
				var depositAmount float64

				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
				accountBalance += depositAmount
				writeBalanceToFile(accountBalance)

				fmt.Println("Updated balance:", accountBalance)
				fmt.Println("------------------------")

			}
		case 3:
			{

				fmt.Print("Withdrawal amount:")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)

				if withdrawAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				if withdrawAmount > accountBalance {
					fmt.Println("You exceeded your balance")
					continue

				} else {

					accountBalance -= withdrawAmount
					writeBalanceToFile(accountBalance)

					fmt.Println("updated balance:", accountBalance)
					fmt.Println("------------------------")
				}

			}
		default:
			{

				fmt.Println("Goodbye!")
				return
			}

		}

	}
}
