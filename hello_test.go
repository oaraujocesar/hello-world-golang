package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		result := Hello("César", "")
		expected := "Hello, César!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Return the default message when no name is provided", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Return the correct greeting when the language is Spanish", func(t *testing.T) {
		result := Hello("César", "Spanish")
		expected := "Hola, César!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Return the correct greeting when the language is French", func(t *testing.T) {
		result := Hello("César", "French")
		expected := "Bonjour, César!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Return the correct greeting when the language is Portuguese", func(t *testing.T) {
		result := Hello("César", "Portuguese")
		expected := "Olá, César!"

		assertCorrectMessage(t, result, expected)
	})
}
