package greetings

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	s := fmt.Sprintf("Hi,%v. Welcome!", name)
	return s
}
