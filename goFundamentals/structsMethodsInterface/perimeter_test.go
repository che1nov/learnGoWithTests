package main

import "testing"

func TestPerimeter(t *testing.T) {
	restangle := Restangle{10.0, 10.0}
	got := Perimeter(restangle)
	want := 40.0

	if got != want {
		t.Errorf("got %2.f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("restangles", func(t *testing.T) {
		restangle := Restangle{12.0, 6.0}
		checkArea(t, restangle, 72.0)
	})

	t.Run("restangles", func(t *testing.T) {
		restangle := Circle{10}
		checkArea(t, restangle, 314.1592653589793)
	})

}

func TestFuncArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Restangle", shape: Restangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
