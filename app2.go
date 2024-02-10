package greetings

import (
	"math/rand"

	"rsc.io/quote"
)

func Hello2(name string) (string, error) {
	message, err := Hello(name)
	message = message + "\n" + quote.Go()
	return message, err
}

func RandomHello(name string) (string, error) {
	// var formats []string
	formats := []string{"Hi, %s", "Howdy, %s", "Wrong un"}
	return Hello(name, formats[rand.Intn(len(formats))])
}
