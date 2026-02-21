package main

import "fmt"

// Constants
const pi = 3.14159
const appName = "Go Deep Foundation"

// Mini project: BMI Calculator
func main() {
	var weight float64
	var height float64

	fmt.Print("Enter weight (kg): ")
	fmt.Scan(&weight)

	fmt.Print("Enter height (meters): ")
	fmt.Scan(&height)

	bmi := weight / (height * height)

	fmt.Println("Your BMI is:", bmi)

	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi < 24.9 {
		fmt.Println("Category: Normal")
	} else if bmi < 29.9 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
}

