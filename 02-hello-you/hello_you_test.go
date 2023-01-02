package main

import "testing"

func TestHelloYou(t *testing.T) {
	got := HelloYou("mateen")
	want := "hello, mateen!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
