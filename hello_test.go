package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris2")
	want := "Hello, Chris2"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
