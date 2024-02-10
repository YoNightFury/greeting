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
	formats := []string{"Hi, %s", "Dowdy, %s"}
	return Hello(name, formats[rand.Intn(len(formats))])
}

func Hellos(names []string) map[string]string {
	hellos := make(map[string]string, len(names))
	for _, name := range names {
		hellos[name], _ = RandomHello(name)
	}
	return hellos

}
