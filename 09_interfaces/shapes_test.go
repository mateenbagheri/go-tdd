package main

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	checkParameter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Parameter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Perimeter of a rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		want := 40.0
		checkParameter(t, rect, want)
	})

	t.Run("Area of a rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rect, want)
	})

	t.Run("Perimeter of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 62.83185307179586
		checkParameter(t, circle, want)
	})

	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		shape         Shape
		wantArea      float64
		wantParameter float64
	}{
		{Rectangle{12, 6}, 72.0, 36.0},
		{Circle{10}, 314.1592653589793, 0},
		{Triangle{10, 8.66025403784, 10, 10, 10}, 43.3012701892, 30},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.wantArea {
			t.Errorf("wanted %g, but instead got %g", tt.wantArea, got)
		}
	}
}
