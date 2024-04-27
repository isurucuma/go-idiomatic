package main

import "fmt"

// const value = 21474836471
const value = 200

func main() {
	// var i int32 = value + 1
	var i int32 = value
	var f = float32(i)
	var fv = value
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(fv)
}
