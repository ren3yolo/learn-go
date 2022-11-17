// declare greetings package
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome !",
		"%v sliding into the matrix",
		"Hail %v, Well met",
	}

	return formats[rand.Intn(len(formats))]
}

// implement package's hello function. A function whose name starts with a capital letter can be called by a function
// not in the same package.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
