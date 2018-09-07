package main

import "testing"

func TestHello(t *testing.T) {
    got := firstgo()
    want := "Hello, My first go app.\n"

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}