// declare greetings package
package greetings

import (
	"errors"
	"fmt"
)

// implement package's hello function. A function whose name starts with a capital letter can be called by a function
// not in the same package.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message, nil
}
