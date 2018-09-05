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

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("got %.2f wanted %.2f", got, tt.hasArea)
			}

		})
	}

	checkArea := func(t *testing.T, shape Shape, wanted float64) {
		t.Helper()
		got := shape.Area()

		if got != wanted {
			t.Errorf("got %.2f, wanted %.2f", got, wanted)
		}
	}

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10.0} // pass in radius
		checkArea(t, circle, 314.1592653589793)
	})
}
