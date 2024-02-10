package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloHappyPath(t *testing.T) {
	name := "Ram"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) {
		t.Fatalf("%q message does not match regex: %#q ", msg, want)
	}
	if err != nil {
		t.Fatalf("Got error: %q", err)
	}
}
func TestHelloUnHappyPath(t *testing.T) {
	name := ""

	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")= %q, %v, want "", error`, msg, err)
	}
}

func TestHellosHappyPath(t *testing.T) {
	names := []string{"Ram", "Shyam", "Jam"}
	template := `\b%s\b`
	msgs := Hellos(names)
	for name, msg := range msgs {
		want := regexp.MustCompile(fmt.Sprintf(template, name))

		if !want.MatchString(msg) {
			t.Fatalf("%q message does not match regex: %#q ", msg, want)
		}
	}
}
