package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greet("Daz")
	want := "Greetings, Daz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
