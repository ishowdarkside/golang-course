package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000
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
