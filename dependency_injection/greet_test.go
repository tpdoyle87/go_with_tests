package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Thomas")

	got := buffer.String()
	want := "Hello, Thomas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
