package main

import "testing"

func TestPerim(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	wanted := 100.0

	if got != wanted {
		t.Errorf("got %.2f, wanted %.2f", got, wanted)
	}
}
