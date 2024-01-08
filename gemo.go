package gemo

import (
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
    // v0.1.5
	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := ""
	return message, nil
}
