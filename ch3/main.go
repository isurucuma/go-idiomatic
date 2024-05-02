package main

import "fmt"

func main() {
	// slice of strings
	var notes = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var n1 = notes[:2]
	var n2 = notes[1:4]
	var n3 = notes[3:]
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	type Employee struct {
		firstname string
		lastname  string
		id        int
	}

	// initialize uding struct literal without names
	var empl1 = Employee{id: 1}
	// initialize using struct literal with names
	var empl2 = Employee{firstname: "Jane", lastname: "Doe", id: 2}

	var empl3 Employee
	empl3.firstname = "John"
	empl3.lastname = "Doe"
	empl3.id = 3

	fmt.Println(empl1)
	fmt.Println(empl2)
	fmt.Println(empl3)
}
