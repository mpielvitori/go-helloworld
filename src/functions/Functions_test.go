package main

import "testing"

func TestFunctions(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Get bigger number from a list", func(t *testing.T) {
		got := GetBiggerNumber(1, 4, 2, 0, 3, 5, 8)
		want := 0
		assertCorrectMessage(t, got, want)
	})
}
