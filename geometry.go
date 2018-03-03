package geometry

import "image"

type Point struct {
	X, Y int
}

type Vector interface {
	GetX() int
	GetY() int
}

type Polygon interface {
	Bounds() []Point
}

type Area interface {
	Origin() Point
	Shape() Polygon
	Map() image.Image
}

type Orientation byte

const (
	Horizontal Orientation = 1 << iota

	Vertical
)

type Occupation byte

const (
	FrontRight Occupation = 1 << iota

	FrontLeft

	RearLeft

	RearRight
)

const (
	FullFront Occupation = FrontLeft | FrontRight

	FullRear Occupation = RearLeft | RearRight

	Full Occupation = FullFront | FullRear
)
