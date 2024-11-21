package fractal

import "testing"

func TestDraw(t *testing.T) {
	err := draw("mandelbrot", make(map[string]string))
	t.Logf("%v", err)
}
