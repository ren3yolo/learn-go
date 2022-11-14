// declare greetings package
package greetings

import "fmt"

// implement package's hello function. A function whose name starts with a capital letter can be called by a function
// not in the same package.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message
}
