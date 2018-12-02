package geometry

import "testing"

func TestPolygonArea(t *testing.T) {
	tables := []struct{
		polygon Polygon
		area float64
	} {
		{SimplePolygon{}, 0},
		{SimplePolygon{Point{0, 0}}, 0},
		{SimplePolygon{Point{0, 0}, Point{0, 0}}, 0},
		{SimplePolygon{Point{0, 0}, Point{0, 0}, Point{0, 0}}, 0},
		{SimplePolygon{Point{0, 0}, Point{0, 0}, Point{0, 0}, Point{0, 0}}, 0},

		// Triangles
		{SimplePolygon{Point{0, 0}, Point{0, 2}, Point{2, 0}}, 2},
		{SimplePolygon{Point{0, 0}, Point{0, 3}, Point{5, 0}}, 7.5},
		{SimplePolygon{Point{1, 2}, Point{5, 2}, Point{3, 5}}, 6},
		{SimplePolygon{Point{-1, 2}, Point{3, 2}, Point{1, 5}}, 6},
		{SimplePolygon{Point{-4, 2}, Point{0, 2}, Point{-2, 5}}, 6},
		{SimplePolygon{Point{-9, 2}, Point{-5, 2}, Point{-7, 5}}, 6},

		// Squares
		{SimplePolygon{Point{0, 0}, Point{0, 2}, Point{2, 0}, Point{2, 2}}, 4},
		{SimplePolygon{Point{-4, 5}, Point{5, 12}, Point{-4, 12}, Point{5, 5}}, 63},

		// Misc
		{SimplePolygon{Point{1.5, 2}, Point{2.2, -7}, Point{6.7, 2}, Point{4.3, -7}}, 32.850000},
	}

	epsilon := 1e-12
	for _, d := range tables {
		if a := PolygonArea(d.polygon); d.area - a > epsilon || a - d.area > epsilon {
			t.Errorf("Expected %f, got %f", d.area, a)
		}
	}
}
