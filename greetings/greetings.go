package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello returns a greeting for the named person.

func Hello(name string) (string, error) {
	// If no name is given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
	// nil means no error, go can return multiple values
}

func Hellos(name []string) (map[string]string, error) {
	// If no name is given, return an error with a message.
	if len(name) == 0 {
		return nil, errors.New("empty name")
	}

	// Create a map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received names and call the Hello function to get the message for each name.
	for _, n := range name {
		message, err := Hello(n)
		if err != nil {
			return nil, err
		}
		messages[n] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages.

func randomFormat() string {
	// A slice of message formats.

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying a random index.
	return formats[rand.Intn(len(formats))]
}
