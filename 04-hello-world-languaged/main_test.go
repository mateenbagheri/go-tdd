package main

import "testing"

func TestHelloWorldLanguaged(t *testing.T) {
	t.Run("testing spanish language", func(t *testing.T) {
		want := "Hola, Mateen"
		got := helloWordLanguaged("Mateen", "Spanish")
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing no name parameter, Spanish", func(t *testing.T) {
		want := "Hola, Mundo"
		got := helloWordLanguaged("", "Spanish")
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing no name parameter, english", func(t *testing.T) {
		want := "Hello, Mateen"
		got := helloWordLanguaged("Mateen", "English")
		assertCorrectMessage(t, want, got)
	})
	t.Run("testing no name parameter, no language parameter", func(t *testing.T) {
		want := "Hello, World"
		got := helloWordLanguaged("", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("the one with random bullshit language input", func(t *testing.T) {
		want := "Hello, Mateen"
		got := helloWordLanguaged("Mateen", "random")
		assertCorrectMessage(t, got, want)
	})

	t.Run("the one with french language input", func(t *testing.T) {
		want := "Bonjour, Mateen"
		got := helloWordLanguaged("Mateen", "French")
		assertCorrectMessage(t, got, want)
	})

	t.Run("the one with french language input and no name parameter", func(t *testing.T) {
		want := "Bonjour, Monde"
		got := helloWordLanguaged("", "French")
		assertCorrectMessage(t, got, want)
	})
}

// assertCorrectMessage simply asserts if a got input is
// equal to a want input.
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q but instead, we got %q", want, got)
	}
}
