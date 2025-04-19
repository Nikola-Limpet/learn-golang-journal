package main

import (
	"fmt"
	"log"

	"go-learn/greetings"
)

func main() {
	// Set properties of the predefined variables
	// the log entry prefix and a flag to diable printing
	// the date and time
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message

	message, err := greetings.Hello("")

	// if an error was returned, print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the message
	// to the console
	fmt.Println(message)
}
