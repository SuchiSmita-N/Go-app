package main

import "testing"

func TestHello(t *testing.T) {

   assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

   t.Run("Hello to people", func(t *testing.T) {	
    got := firstgo("suchi")
    want := "Hello, suchi"
    assertCorrectMessage(t, got, want)
    })

   t.Run("empty string: World", func(t *testing.T) {
      got := firstgo("")
      want := "Hello, World"
      assertCorrectMessage(t, got, want)   
    })
}
