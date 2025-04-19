package greetings

import (
	"errors"
	"fmt"
)

// hello returns a greeting for the named person.

func Hello(name string) (string, error) {
	// If no name is given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
	// nil means no error, go can return multiple values
}
