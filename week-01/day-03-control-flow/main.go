package main

import "fmt"

func main() {
	age := 22
	balance := 1200

	// IF ELSE
	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}

	// IF ELSE IF
	if balance < 500 {
		fmt.Println("Low balance")
	} else if balance < 1000 {
		fmt.Println("Medium balance")
	} else {
		fmt.Println("Sufficient balance")
	}

	// SWITCH
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}
}
