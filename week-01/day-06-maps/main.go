package main

import "fmt"

func main() {

    // map of student and grades
    grades := map[string]string{
        "Rahul":  "A",
        "Ganesh": "A+",
        "Riya":   "B",
        "Aman":   "A",
    }

    fmt.Println("Student Grades:")

    for name, grade := range grades {
        fmt.Println(name, "->", grade)
    }

    // check specific student
    fmt.Println("Ganesh grade:", grades["Ganesh"])
}