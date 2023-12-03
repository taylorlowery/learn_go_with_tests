package main

import "testing"

func TestHello(t *testing.T) {
	got := greet("Tater")
	want := "Howdy, Tater!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
