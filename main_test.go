package main

import "testing"

func TestGetString(t *testing.T) {

	got := GetString()
	want := "hello world"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
