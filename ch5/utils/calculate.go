package utils

import "errors"

var operations = map[string]func(int, int) (int, error){
	"+": func(a, b int) (int, error) { return a + b, nil },
	"-": func(a, b int) (int, error) { return a - b, nil },
	"*": func(a, b int) (int, error) { return a * b, nil },
	"/": func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	},
}

func calculate(a int, b int, op string) (int, error) {
	if operation, ok := operations[op]; ok {
		return operation(a, b)
	}
	return 0, errors.New("operation not found")
}

func Run(a int, b int, op string) {
	result, err := calculate(a, b, op)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
}
