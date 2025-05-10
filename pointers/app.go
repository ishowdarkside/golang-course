package main

import "fmt"

func main() {

	age := 32 // Regular Variable
	agePointer := &age

	fmt.Println("Age prije getAdultYears poziva:", age)
	getAdultYears(agePointer)
	fmt.Println("Age poslije getAdultYears poziva:", age)

}

func getAdultYears(age *int) {

	*age = *age - 18

}
