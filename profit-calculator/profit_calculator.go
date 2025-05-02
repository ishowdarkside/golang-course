package main

import "fmt"

func main() {


	revenue := readUserInput("Revenue")
	expenses := readUserInput("Expenses")	
	taxRate := readUserInput("Tax Rate")

	ebt, profit, ratio := calculateIncomes(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}


func readUserInput(text string) float64 {

	fmt.Print(text, ":")
	var userInput float64
   	fmt.Scan(&userInput)

   return userInput

}


func calculateIncomes(revenue float64, expenses float64,taxRate float64) (float64, float64, float64) {


	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio:= ebt / profit


	return ebt,profit,ratio


}