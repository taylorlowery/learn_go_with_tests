package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Greet("Tater", "")
		want := "Howdy, Tater!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is provided", func(t *testing.T) {
		got := Greet("", "")
		want := "Howdy, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Greet("Tater", "Spanish")
		want := "Hola, Tater!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Greet("Tater", "English")
		want := "Hello, Tater!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Greet("Tater", "French")
		want := "Bonjour, Tater!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
