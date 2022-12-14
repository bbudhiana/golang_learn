package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with name, checking for valid return value
func TestHelloName(t *testing.T) {
	name := "Bana"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Bana")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Bana") = %q, %v, want match for %#q, nill`, msg, err, want)
	}
}

// TestHelloEmpty calls greeting.Hello with an empty string, checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
