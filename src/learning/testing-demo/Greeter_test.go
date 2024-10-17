package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMorningGreet(t *testing.T) {
	got := MorningGreet("Jacob")
	want := "Good morning, Jacob! How are you doing today?"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
