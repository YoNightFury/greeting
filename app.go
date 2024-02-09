package greetings

import "fmt"

func Hello(name string) string {
	// longer way of declaring variables
	// var message string
	// message = fmt.Sprintf("Hello, %s", name)

	// shortcut variable declaration and initialization
	message := fmt.Sprintf("Hello, %s", name)
	return message
}
