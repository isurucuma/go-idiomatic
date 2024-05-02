package main

import (
	"fmt"
	"functions/utils"
)

func main() {
	a, b := 10, 2
	op := "/"
	utils.Run(a, b, op)

	// file length
	filename := "main.go"
	length, err := utils.Filelen(filename)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Length:", length, "bytes")
	}

	// prefix
	helloPrefix := utils.Prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
