package main

import "testing"

func TestLang(t *testing.T) {

   assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }
   t.Run("in hindi", func(t *testing.T) {
        got := firstgo("suchi", "Hindi")
        want := "नमस्ते, suchi"
        assertCorrectMessage(t, got, want)
    })
}

