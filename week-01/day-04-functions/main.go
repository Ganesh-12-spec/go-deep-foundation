package main

import "fmt"

/*
DAY 4: FUNCTIONS
*/

// 1. Simple function
func greet() {
	fmt.Println("Hello from Go function")
}

// 2. Function with parameters
func add(a int, b int) int {
	return a + b
}

// 3. Multiple return values
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

// 4. Named return values
func rectangleArea(length int, width int) (area int) {
	area = length * width
	return // naked return
}

// 5. Variadic function (many inputs)
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {

	// Call simple function
	greet()

	// Call add
	result := add(10, 20)
	fmt.Println("Add:", result)

	// Call divide
	q, r := divide(10, 3)
	fmt.Println("Divide:", q, r)

	// Call named return function
	area := rectangleArea(5, 4)
	fmt.Println("Area:", area)

	// Call variadic function
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sum)

	fmt.Println("Subtract:", subtract(10, 5))
  fmt.Println("Multiply:", multiply(4, 6))
}