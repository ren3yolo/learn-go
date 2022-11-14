package main

import (
	"fmt"

	"learning.com/greetings"
)

func main() {
	message := greetings.Hello("Raj")
	fmt.Println(message)
}
