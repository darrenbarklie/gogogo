package main

import "testing"

func TestHelloGreet(t *testing.T) {
	t.Run("default 'Hello, world' when empty string supplied", func(t *testing.T) {
		got := HelloGreet("", "")
		want := "Hello, world"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("greeting people by name", func(t *testing.T) {
		got := HelloGreet("Daz", "")
		want := "Hello, Daz"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("in Vietnamese language", func(t *testing.T) {
		got := HelloGreet("Daz", "Vietnamese")
		want := "Chao, Daz"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("in French language", func(t *testing.T) {
		got := HelloGreet("Daz", "French")
		want := "Bonjour, Daz"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
