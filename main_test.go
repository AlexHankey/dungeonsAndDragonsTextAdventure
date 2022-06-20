package main

import (
	"testing"
)

func Hello(t *testing.T) {
	got := hello()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
