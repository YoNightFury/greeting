package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Hello2(name string) string {
	// longer way of declaring variables
	// var message string
	// message = fmt.Sprintf("Hello, %s", name)

	// shortcut variable declaration and initialization
	message := fmt.Sprintf("Hello, %s", name)
	message = message + "\n" + quote.Go()
	return message
}
