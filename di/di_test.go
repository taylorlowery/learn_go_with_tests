package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("initial test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Tater")
		got := buffer.String()
		want := "Howdy, Tater!"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
