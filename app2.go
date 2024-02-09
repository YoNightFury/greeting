package greetings

import (
	"rsc.io/quote"
)

func Hello2(name string) string {
	message := Hello(name)
	message = message + "\n" + quote.Go()
	return message
}
