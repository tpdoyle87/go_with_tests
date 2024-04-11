package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 3.0, Height: 3.0}
	got := Perimeter(rectangle)
	want := 12.0

	assertCorrectMessage(t, got, want, rectangle)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 4.0, Height: 4.0}, 16.0},
		{Circle{Radius: 4}, 50.26548245743669},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertCorrectMessage(t, got, tt.want, tt.shape)
	}
}

func assertCorrectMessage(t testing.TB, got, want float64, shape Shape) {
	t.Helper()
	if got != want {
		t.Errorf("%#v %g but got %g", shape, want, got)
	}
}
