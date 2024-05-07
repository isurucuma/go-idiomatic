package main

import (
	"fmt"
	"generics/lists"
	"generics/utils"
)

// Question 1 -----------------------------------------------------------------------------------------------

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func double[T Number](x T) T {
	return x + x
}

func Convert[T1, T2 Number](in T1) T2 {
	return T2(in)
}

// Question 2 -----------------------------------------------------------------------------------------------
type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type PrintableInt int

func (p PrintableInt) String() string {
	return fmt.Sprintf("printing int : %d", p)
}

type PrintableFloat float64

func (p PrintableFloat) String() string {
	return fmt.Sprintf("printint float : %f", p)
}

func printNum[T Printable](num T) {
	fmt.Println(num.String())
}

// here the important concept is that Printable is a type constraint that can only be used to constrain the type of the generic type parameter
// and not the type of the value that is passed to the function or a type of a variable.

// Question 3 -----------------------------------------------------------------------------------------------

func main() {
	fmt.Println(double(2))
	fmt.Println(double(3.14))

	a := 10
	b := Convert[int, int64](a)
	fmt.Println(b)
	fmt.Println(Convert[int, float64](42))

	var pint PrintableInt = 4
	var pfloat PrintableFloat = 3.14

	printNum(pint)
	printNum(pfloat)

	fmt.Println("Running utils main----------------------------------------------------------------------")

	utils.UtilsMain()

	fmt.Println("Running linked list main----------------------------------------------------------------------")

	lists.TestLinkedList()

}
