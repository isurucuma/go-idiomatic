package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed data/english_rights.txt
var englishRights string

//go:embed data/arabic_rights.txt
var arabicRights string

//go:embed data/sinhala_rights.txt
var sinhalaRights string

func main() {
	if len(os.Args) > 1 {
		language := os.Args[1]
		switch language {
		case "english":
			fmt.Println(englishRights)
		case "arabic":
			fmt.Println(arabicRights)
		case "sinhala":
			fmt.Println(sinhalaRights)
		default:
			fmt.Println("Invalid language")
		}
	}
}
