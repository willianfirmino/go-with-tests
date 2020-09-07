package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("William")
	want := "Hello, William"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
