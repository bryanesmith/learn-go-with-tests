package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Villager")
	want := "Hello, Villager"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
