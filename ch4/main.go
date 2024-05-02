package main

import (
	"fmt"
	"math/rand"
)

func main() {
	vals := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		vals = append(vals, rand.Intn(100))
	}

	for i, v := range vals {
		println(i, v)
		// this is not the most idiomatic way to do this, instead use if else
		// switch {
		// case (v%2 == 0) && (v%3 == 0):
		// 	fmt.Println("Six!")

		// case v%2 == 0:
		// 	fmt.Println("Two!")

		// case v%3 == 0:
		// 	fmt.Println("Three!")

		// default:
		// 	fmt.Println("Never mind")
		// }

		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
			continue
		}

		if v%2 == 0 {
			fmt.Println("Two!")
			continue
		}

		if v%3 == 0 {
			fmt.Println("Three!")
			continue
		}

		fmt.Println("Never mind")
	}

	fmt.Println("Printing the total")
	// printing 0-10 total
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println("Total:", total)
	// result is 0 which is not the expected result. That is because of the shadowing of the total variable
}
