package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	s := greetings.Hello("Gladys")
	fmt.Println(s)
}
