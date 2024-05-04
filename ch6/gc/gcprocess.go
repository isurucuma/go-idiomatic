package gc

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func TestGc() {
	const size = 10000000
	// var persons [size]Person
	var persons = make([]Person, 0, size)
	for i := 0; i < size; i++ {
		// persons[i] = Person{Firstname: "James", Lastname: "Bond", Age: 40}
		persons = append(persons, Person{Firstname: "James", Lastname: "Bond", Age: 40})
	}
	fmt.Println(persons[size-1])
}
