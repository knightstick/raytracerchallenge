package tuples

import "math"

// A Tuplelike can be used like a Tuple, for things like the Dot product
type Tuplelike interface {
	Values() []float64
	At(idx int) float64
}

// A Pointlike translates values to coordinates, like X and Y
type Pointlike interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

// A Tuple is a list of float64s of any length
type Tuple []float64

// Values returns the contents of the Tuple
func (t Tuple) Values() []float64 {
	return t
}

// At returns a value at a particular index
func (t Tuple) At(idx int) float64 {
	return t[idx]
}

// A Point represents a point in 3 dimensional space
type Point []float64

// Values returns the contents as a Tuplelike list
func (p Point) Values() []float64 {
	return p
}

// At returns the value at a particular index
func (p Point) At(idx int) float64 {
	return p[idx]
}

// X is a cardinal direction
func (p Point) X() float64 {
	return p[0]
}

// Y is a cardinal direction
func (p Point) Y() float64 {
	return p[1]
}

// Z is a cardinal direction
func (p Point) Z() float64 {
	return p[2]
}

// W can be used to discriminate a Point Tuple from a Vector one
func (p Point) W() float64 {
	return p[3]
}

// A Vector represents a quantity with a direction in 3 dimensions
type Vector []float64

// Values returns the contents as a Tuplelike list
func (v Vector) Values() []float64 {
	return v
}

// At returns the value at a particular index
func (v Vector) At(idx int) float64 {
	return v[idx]
}

// X is a cardinal direction
func (v Vector) X() float64 {
	return v[0]
}

// Y is a cardinal direction
func (v Vector) Y() float64 {
	return v[1]
}

// Z is a cardinal direction
func (v Vector) Z() float64 {
	return v[2]
}

// W can be used to discriminate a Point Tuple from a Vector one
func (v Vector) W() float64 {
	return v[3]
}

// Color is a length 3 Tuple containing Red, Green and Blue components
type Color []float64

// Values returns the contents as a Tuplelike list
func (c Color) Values() []float64 {
	return c
}

// At returns the value at a particular index
func (c Color) At(idx int) float64 {
	return c[idx]
}

// Red returns the red component of the color
func (c Color) Red() float64 {
	return c[0]
}

// Green returns the green component of the color
func (c Color) Green() float64 {
	return c[1]
}

// Blue returns the blue component of the color
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

// NewColor instantiates a new Color
func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

// IsPoint can be used to check if a Tuplelike is a Point or not
func IsPoint(t Tuplelike) bool {
	return t.At(3) == pointW
}

// IsVector can be used to check if a Tuplelike is a Vector or not
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

// MultiplyColors takes two colors and returns their hadamard_product
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
