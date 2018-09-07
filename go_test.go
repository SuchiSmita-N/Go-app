package main

import "testing"

func TestHello(t *testing.T) {
    got := firstgo("suchi.")
    want := "Hello, suchi."

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}