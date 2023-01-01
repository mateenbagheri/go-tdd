package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Mateen has started learning more about TDD!"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
