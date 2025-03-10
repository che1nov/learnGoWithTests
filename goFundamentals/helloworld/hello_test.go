package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := HelloWorld("Ilya", "English")
		want := "Hello, Ilya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, void!' when an empty string is supplied", func(t *testing.T) {
		got := HelloWorld("", "English")
		want := "Hello, void"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloWorld("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Franch", func(t *testing.T) {
		got := HelloWorld("Lui", "French")
		want := "Bonjour, Lui"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	//t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
