package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("to a person", func(t *testing.T) {
		got := Hello()
		want := "Hello Martin!"
		assertCorrectMessage(t, got, want)
	})
}
