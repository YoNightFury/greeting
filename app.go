package greetings

import (
	"errors"
	"fmt"
	"strings"
)

func Hello(name string, format ...string) (string, error) {

	if name == "" {
		return "", errors.New("name is mandatory")
	}

	// longer way of declaring variables
	// var message string
	// message = fmt.Sprintf("Hello, %s", name)

	if len(format) > 0 {
		if strings.Contains(format[0], "%s") {
			return fmt.Sprintf(format[0], name), nil
		}
		return "", errors.New("the formated string should contain %s")
	}
	// shortcut variable declaration and initialization
	message := fmt.Sprintf("Hello, %s", name)
	return message, nil
}
