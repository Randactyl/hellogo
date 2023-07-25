package main

import (
	"fmt"

	"Randactyl/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Shirley")
	fmt.Println(message)
}
