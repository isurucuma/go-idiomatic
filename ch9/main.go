package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Exercise 9.1 Sentinel error ----------------------------------------------------------------------

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrInvalidId = Sentinel("invalid ID")
	ErrMissingId = Sentinel("missing ID")
)

// Exercise 9.2 Custom error type -------------------------------------------------------------------

type EmptyFieldError struct {
	Field string
}

func (efe EmptyFieldError) Error() string {
	return fmt.Sprintf("missing %s", efe.Field)
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		if err != nil {
			if errors.Is(err, ErrInvalidId) {
				fmt.Println(err)
			}
			if errors.As(err, &EmptyFieldError{}) {
				fmt.Println(err)
			}
			// fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	var errs []error
	if len(e.ID) == 0 {
		errs = append(errs, ErrMissingId)
	}
	if !validID.MatchString(e.ID) {
		errs = append(errs, ErrInvalidId)
	}
	if len(e.FirstName) == 0 {
		errs = append(errs, EmptyFieldError{"FirstName"})
	}
	if len(e.LastName) == 0 {
		errs = append(errs, EmptyFieldError{"LastName"})
	}
	if len(e.Title) == 0 {
		errs = append(errs, EmptyFieldError{"Title"})
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}
