package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Thomas")
    want := "Hello, Thomas"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}