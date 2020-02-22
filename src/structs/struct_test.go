package structs

import "testing"

func TestPerimeter(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	}
	rec := Rectangle{10.0, 10.0}
	got := rec.Perimeter()
	want := 40.0
	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %g but got %g", want, got)
		}
	}
	t.Run("Testing Area's Function for Rectangle", func(t *testing.T) {
		rec := Rectangle{10.0, 10.0}
		checkArea(t, rec, 100.0)
	})
	t.Run("Testing Area's Function for Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 31.400000000000002)
	})
}
