package main

import "testing"

func TestFunctionsNested(t *testing.T) {
	t.Run("Count down from 5 with multiple rockets", func(t *testing.T) {
		err := CountDownMultiple(5, "R1", "Rocket2")
		if err != nil {
			t.Errorf("got error %q", err)
		}
	})
}
