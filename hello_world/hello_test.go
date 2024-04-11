package main

import "testing"

func TestHelloName(t *testing.T) {
    got := Hello("Thomas", "")
    want := "Hello, Thomas"

    assertCorrectMessage(t, got, want)
}

func TestHelloWorld(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"

    assertCorrectMessage(t, got, want)
}

func TestHelloSpanish(t *testing.T) {
    got := Hello("Thomas", "Spanish")
    want := "Hola, Thomas"

    assertCorrectMessage(t, got, want)
}

func TestHelloFrench(t *testing.T) {
    got := Hello("Thomas", "French")
    want := "Bonjour, Thomas"

    assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

