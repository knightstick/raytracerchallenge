package tuples

import "math"

type Tuple []float64

type Point []float64

type Pointlike interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

func (p Point) X() float64 {
	return p[0]
}

func (p Point) Y() float64 {
	return p[1]
}

func (p Point) Z() float64 {
	return p[2]
}

func (p Point) W() float64 {
	return p[3]
}

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

const pointW = 1.0
const vectorW = 0.0

// New instantiates a new Tuple
func New(args ...float64) Tuple {
	return args
}

// NewPoint instatiates a new Point
func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, pointW}
}

// NewVector instatiates a new Vector
func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, vectorW}
}

func NewColor(r, g, b float64) Color {
	return Color{Red: r, Green: g, Blue: b}
}

// IsPoint checks if a Tuple is a Point
func (t Tuple) IsPoint() bool {
	// Can this crash?
	return t[3] == pointW
}

// IsVector checks if a Tuple is a Vector
func (t Tuple) IsVector() bool {
	return t[3] == vectorW
}

// Equal checks all fields in the Tuple are within a small distance from each
// other
func Equal(t, other Tuple) bool {
	for idx, value := range t {
		if !inEpsilon(value, other[idx]) {
			return false
		}

	}
	return true
}

// Add returns a new Tuple which results from adding the other Tuple to the
// first
func Add(t, other Tuple) Tuple {
	result := []float64{}

	for idx, value := range t {
		result = append(result, value+other[idx])
	}

	return result
}

// Subtract returns a new Tuple which results from subtracting the other Tuple to the
// first
func Subtract(t, other Tuple) Tuple {
	result := []float64{}

	for idx, value := range t {
		result = append(result, value-other[idx])
	}

	return result
}

// Negate returns a new Tuple which negates each component of the Tuple
func Negate(t Tuple) Tuple {
	result := []float64{}

	for _, value := range t {
		result = append(result, -value)
	}

	return result
}

// Multiply returns a new Tuple which represents the Tuple multiplied by a
// scalar value
func Multiply(t Tuple, n float64) Tuple {
	result := []float64{}

	for _, value := range t {
		result = append(result, value*n)
	}

	return result
}

// Divide returns a new Tuple which represents the Tuple divided by a scalar
// value
func Divide(t Tuple, n float64) Tuple {
	return Multiply(t, 1/n)
}

// Normalize returns a unit Vector in the same direction as the Tuple
func Normalize(t Tuple) Tuple {
	return Divide(t, Magnitude(t))
}

// Magnitude returns the length of a Vector, computed by summing the squares
// of each value, then taking the square root
func Magnitude(t Tuple) float64 {
	sum := 0.0

	for _, value := range t {
		sum += value * value
	}

	return math.Sqrt(sum)
}

// Dot or scalar product returns a scalar value from two Vectors, can be used
// to represent the angle between two Vectors
func Dot(p, other Pointlike) float64 {
	return (p.X() * other.X()) +
		(p.Y() * other.Y()) +
		(p.Z() * other.Z()) +
		(p.W() * other.W())
}

// Cross or Vector product returns a new Vector that is perpendicular to both
// Vectors
func Cross(p, other Pointlike) Tuple {
	return NewVector((p.Y()*other.Z())-(p.Z()*other.Y()),
		(p.Z()*other.X())-(p.X()*other.Z()),
		(p.X()*other.Y())-(p.Y()*other.X()))
}

const epsilon = 0.0001

func inEpsilon(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
