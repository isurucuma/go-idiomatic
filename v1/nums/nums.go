package nums

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// # Add <- this is a comment header
//
// This is a link to get more info [mathsisfun].
//
// [mathsisfun]: https://www.mathsisfun.com/numbers/addition.html
// Add returns the sum of two integers
func Add[T Number](a, b T) T {
	return a + b
}
