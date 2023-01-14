package main

import "testing"

func TestMultipleSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := multiple_sum(numbers)
	want := 15
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d but got number %d instead", want, got)
	}
}
