package main

import "testing"

func TestGetPong(t *testing.T) {
	want := "pong"
	got := getPong()
	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
