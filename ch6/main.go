package main

import (
	"fmt"
	"pointers/gc"
)

func makePerson(firstname, lastname string, age int) gc.Person {
	return gc.Person{Firstname: firstname, Lastname: lastname, Age: age}
}

func makePersonPointer(firstname, lastname string, age int) *gc.Person {
	return &gc.Person{Firstname: firstname, Lastname: lastname, Age: age}
}

func updateSlice(s []string, v string) {
	s[len(s)-1] = v
	fmt.Println("Slice in updateSlice", s)
}

func growSlice(s []string, v string) {
	s = append(s, v)
	fmt.Println("Slice in growSlice", s)
}

func main() {
	makePerson("James", "Bond", 40)
	makePersonPointer("Miss", "Moneypenny", 30)
	// fmt.Println("Person 1", p1)
	// fmt.Println("Person 2", p2)

	s := []string{"a", "b", "c"}
	fmt.Println("Slice before update", s)
	updateSlice(s, "d")
	fmt.Println("Slice after udpate", s)

	s = []string{"a", "b", "c"}
	fmt.Println("Slice before grow", s)
	growSlice(s, "d")
	fmt.Println("Slice after grow", s)

	gc.TestGc()
}
