package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ethan", "english")
		want := "Hello, Ethan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is applied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Maria", "spanish")
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Francis", "french")
		want := "Bonjour, Francis"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
