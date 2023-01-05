package main

import "testing"

func TestHelloWorldAgain(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := helloWorldAgain("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to the world", func(t *testing.T) {
		got := helloWorldAgain("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
