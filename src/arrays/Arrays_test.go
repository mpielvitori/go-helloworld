package main

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
	})

	t.Run("appending slices", func(t *testing.T) {

		firstSlice := []int{1, 2, 3}
		secondSlice := []int{4, 5, 6}

		got := Sum(append(firstSlice, secondSlice...))
		want := 21

		assertCorrectMessage(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertCorrectMessage(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrectMessage(t, got, want)
	})

}
