package main

import "testing"

func TestMultipleSum(t *testing.T) {
	t.Run("testing the function multiple_sum by giving array to it", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		want := 15
		assertCorrectMessageSum(t, got, want)
	})
}

func assertCorrectMessageSum(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d but got number %d instead", want, got)
	}
}
