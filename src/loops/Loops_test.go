package main

import "testing"

func TestFunctions(t *testing.T) {
	t.Run("Count down from 5", func(t *testing.T) {
		err := CountDown(5)
		if err != nil {
			t.Errorf("got error %q", err)
		}
	})
}
