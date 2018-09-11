package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"hola": "hello"}

	got := Search(dictionary, "hola")
	want := "hello"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
