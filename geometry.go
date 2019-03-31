package geometry

import (
	"sort"
)

/**
 * -----
 * Point
 * -----
 */

type Point struct {
	X, Y float64
}

func (p Point) Start() Point {
	return Point{0, 0}
}

func (p Point) End() Point {
	return p
}

/**
 * ------
 * Vector
 * ------
 */

type Vector interface {
	Start() Point
	End() Point
}

type vector [2]Point

func (v vector) Start() Point {
	return v[0]
}

func (v vector) End() Point {
	return v[1]
}

func VectorLength(v Vector) float64 {
	// TODO: Implement

	return 0
}

func VectorAngle(v Vector) float64 {
	// TODO: Implement

	return 0
}

func AddVectors(vs ... Vector) Vector {
	// TODO: Implement

	return nil
}

func ReverseVector(v Vector) Vector {
	return vector{
		v.End(),
		v.Start(),
	}
}

/**
 * -------
 * Polygon
 * -------
 */

type Polygon interface {
	Coords() []Point
}

func PolygonLen(p Polygon) int {
	return len(p.Coords())
}

func PolygonCenter(p Polygon) Point {
	pts := p.Coords()

	c := Point{0, 0}
	for _, pt := range pts {
		c.X += pt.X
		c.Y += pt.Y
	}

	c.X /= float64(len(pts))
	c.Y /= float64(len(pts))

	return c
}

// Smallest Enclosing Rectangle (Minimum Bounding Box)
func PolygonEnvelope(p Polygon) Polygon {
	// TODO: Implement

	return nil
}

func PolygonArea(p Polygon) float64 {
	// Check if p is not actually a point or line
	if PolygonLen(p) < 3 {
		return 0
	}

	sp := ClockwiseSortedPolygon(p)
	pts := sp.Coords()
	dArea := 0.0
	for j, i := len(pts)-1, 0; i < len(pts); j, i = i, i+1 {
		dArea += (pts[j].X + pts[i].X) * (pts[j].Y - pts[i].Y)
	}

	return dArea / 2
}

/**
 * --------------
 * Simple Polygon
 * --------------
 */

type SimplePolygon []Point

func (sp SimplePolygon) Coords() []Point {
	return []Point(sp)
}

/**
 * ---------
 * Rectangle
 * ---------
 */

 // TODO: Implement a rotated rectangle

/**
 * ----------------
 * Sortable Polygon
 * ----------------
 */

type SortablePolygon struct {
	coords []Point
	center Point
}

func NewSortablePolygon(p Polygon) SortablePolygon {
	return SortablePolygon{p.Coords(), PolygonCenter(p)}
}

func (sp SortablePolygon) Coords() []Point {
	return sp.coords
}

func (sp SortablePolygon) Center() Point {
	return sp.center
}

func (sp SortablePolygon) Len() int {
	return len(sp.coords)
}

func (sp SortablePolygon) Less(i, j int) bool {
	a, b, center := sp.coords[i], sp.coords[j], sp.center

	if a.X-center.X >= 0 && b.X-center.X < 0 {
		return true
	}

	if a.X-center.X < 0 && b.X-center.X >= 0 {
		return false
	}

	if a.X-center.X == 0 && b.X-center.X == 0 {
		if a.Y-center.Y >= 0 || b.Y-center.Y >= 0 {
			return a.Y > b.Y
		}

		return b.Y > a.Y
	}

	// compute the cross product of vectors (center -> a) x (center -> b)
	if det := (a.X-center.X)*(b.Y-center.Y) - (b.X-center.X)*(a.Y-center.Y); 0 != det {
		return det < 0
	}

	// points a and b are on the same line from the center
	// check which point is closer to the center
	d1 := (a.X-center.X)*(a.X-center.X) + (a.Y-center.Y)*(a.Y-center.Y)
	d2 := (b.X-center.X)*(b.X-center.X) + (b.Y-center.Y)*(b.Y-center.Y)

	return d1 > d2
}

func (sp *SortablePolygon) Swap(i, j int) {
	sp.coords[i], sp.coords[j] = sp.coords[j], sp.coords[i]
}

func ClockwiseSortedPolygon(p Polygon) Polygon {
	sp := NewSortablePolygon(p)
	sort.Sort(&sp)

	return sp
}

func CounterClockwiseSortedPolygon(p Polygon) Polygon {
	sp := NewSortablePolygon(p)
	sort.Sort(sort.Reverse(&sp))

	return sp
}
