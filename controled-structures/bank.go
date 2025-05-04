package main

import (
	"bankingApp/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	accountBalance, err := fileops.GetFloatFromFile("account-balance.txt", 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		//panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Reach us 24/7: ", randomdata.PhoneNumber())

	for {

		PresentOptions()

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
				fileops.WriteFloatToFile(accountBalance, "account-balance.txt")

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
					fileops.WriteFloatToFile(accountBalance, "account-balance.txt")

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
