package main

import "fmt"

func main() {

    // array of marks
    marks := [5]int{78, 85, 90, 67, 88}

    total := 0

    for i := 0; i < len(marks); i++ {
        total += marks[i]
    }

    average := total / len(marks)

    fmt.Println("Marks:", marks)
    fmt.Println("Total:", total)
    fmt.Println("Average:", average)

    // slice example
    topStudents := marks[1:4]

    fmt.Println("Top students slice:", topStudents)
}