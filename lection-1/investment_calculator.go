package main

import (
	"fmt"
	"math"
)

func main() {
	
	const inflationRate float64 = 2.5;


	var investmentAmount  float64
	var years float64
	var expectedReturnRate float64


	fmt.Print("Investement amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Println();

	fmt.Print("Years:");
	fmt.Scan(&years)


	fmt.Println()

	fmt.Print("Expected return rate:")
	fmt.Scan(&expectedReturnRate)


	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureRealValue)
	

}