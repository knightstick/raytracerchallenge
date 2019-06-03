package tuples

import "math"

type Tuplelike interface {
	Values() []float64
	At(idx int) float64
}

type Pointlike interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

type Tuple []float64

func (t Tuple) Values() []float64 {
	return t
}

func (t Tuple) At(idx int) float64 {
	return t[idx]
}

type Point []float64

func (p Point) Values() []float64 {
	return p
}

func (p Point) At(idx int) float64 {
	return p[idx]
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

type Vector []float64

func (v Vector) Values() []float64 {
	return v
}

func (v Vector) At(idx int) float64 {
	return v[idx]
}

func (v Vector) X() float64 {
	return v[0]
}

func (v Vector) Y() float64 {
	return v[1]
}

func (v Vector) Z() float64 {
	return v[2]
}

func (v Vector) W() float64 {
	return v[3]
}

type Color []float64

func (c Color) Values() []float64 {
	return c
}

func (c Color) At(idx int) float64 {
	return c[idx]
}

func (c Color) Red() float64 {
	return c[0]
}

func (c Color) Green() float64 {
	return c[1]
}

func (c Color) Blue() float64 {
	return c[2]
}

const pointW = 1.0
const vectorW = 0.0

// New instantiates a new Tuple
func New(args ...float64) Tuple {
	return args
}

// NewPoint instatiates a new Point
func NewPoint(x, y, z float64) Point {
	return Point{x, y, z, pointW}
}

// NewVector instatiates a new Vector
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z, vectorW}
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func IsPoint(t Tuplelike) bool {
	return t.At(3) == pointW
}

func IsVector(t Tuplelike) bool {
	return t.At(3) == vectorW
}

// Equal checks all fields in the Tuple are within a small distance from each
// other
func Equal(t, other Tuplelike) bool {
	for idx, value := range t.Values() {
		if !inEpsilon(value, other.At(idx)) {
			return false
		}

	}
	return true
}

// Add returns a new Tuple which results from adding the other Tuple to the
// first
func Add(t, other Tuplelike) Tuplelike {
	result := []float64{}

	for idx, value := range t.Values() {
		result = append(result, value+other.At(idx))
	}

	return Tuple(result)
}

// Subtract returns a new Tuple which results from subtracting the other Tuple to the
// first
func Subtract(t, other Tuplelike) Tuplelike {
	result := []float64{}

	for idx, value := range t.Values() {
		result = append(result, value-other.At(idx))
	}

	return Tuple(result)
}

// Negate returns a new Tuple which negates each component of the Tuple
func Negate(t Tuplelike) Tuplelike {
	result := []float64{}

	for _, value := range t.Values() {
		result = append(result, -value)
	}

	return Tuple(result)
}

// Multiply returns a new Tuple which represents the Tuple multiplied by a
// scalar value
func Multiply(t Tuplelike, n float64) Tuplelike {
	result := []float64{}

	for _, value := range t.Values() {
		result = append(result, value*n)
	}

	return Tuple(result)
}

func MultiplyColors(c1, c2 Color) Color {
	return NewColor(c1.Red()*c2.Red(),
		c1.Green()*c2.Green(),
		c1.Blue()*c2.Blue())
}

// Divide returns a new Tuple which represents the Tuple divided by a scalar
// value
func Divide(t Tuplelike, n float64) Tuplelike {
	return Multiply(t, 1/n)
}

// Normalize returns a unit Vector in the same direction as the Tuple
func Normalize(t Tuplelike) Tuplelike {
	return Divide(t, Magnitude(t))
}

// Magnitude returns the length of a Vector, computed by summing the squares
// of each value, then taking the square root
func Magnitude(t Tuplelike) float64 {
	sum := 0.0

	for _, value := range t.Values() {
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
	return Tuple(NewVector((p.Y()*other.Z())-(p.Z()*other.Y()),
		(p.Z()*other.X())-(p.X()*other.Z()),
		(p.X()*other.Y())-(p.Y()*other.X())))
}

const epsilon = 0.0001

func inEpsilon(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
