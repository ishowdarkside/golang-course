package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, revenueErr := readUserInput("Revenue")
	expenses, expensesErr := readUserInput("Expenses")
	taxRate, taxRateErr := readUserInput("Tax Rate")

	if revenueErr != nil || expensesErr != nil || taxRateErr != nil {

		fmt.Println("Negative values used. Exiting application")
		return

	}
	ebt, profit, ratio := calculateIncomes(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

	writeValuesToFile(ebt, profit, ratio)
}

func readUserInput(text string) (float64, error) {

	fmt.Print(text, ":")
	var userInput float64
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("invalid value. Cannot be <= 0")
	}
	return userInput, nil

}

func calculateIncomes(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio

}

func writeValuesToFile(ebt float64, profit float64, ratio float64) {

	computedValues := fmt.Sprintf("EBT: %.1f, PROFIT: %.1f, RATIO: %.1f", ebt, profit, ratio)
	os.WriteFile("output.txt", []byte(computedValues), 0644)

}
