package geometry

type Point struct {
	X, Y float64
}

type Vector interface {
	GetX() float64
	GetY() float64
}

type Polygon interface {
	Coords() []Point
}

type Area interface {
	Origin() Point
	Shape() Polygon
	MapFile() string
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
