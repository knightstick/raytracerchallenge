package tuples

import "math"

// A Tuple is a list of 4 floats used to represent Points and Vectors
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

const pointW = 1.0
const vectorW = 0.0

// NewTuple instantiates a new Tuple
func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

// NewPoint instatiates a new Point
func NewPoint(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: pointW}
}

// NewVector instatiates a new Vector
func NewVector(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: vectorW}
}

// IsPoint checks if a Tuple is a Point
func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector checks if a Tuple is a Vector
func (t Tuple) IsVector() bool {
	return t.W == 0
}

// Equal checks all fields in the Tuple are within a small distance from each
// other
func Equal(t, other Tuple) bool {
	return inEpsilon(t.X, other.X) && inEpsilon(t.Y, other.Y) &&
		inEpsilon(t.Z, other.Z) && inEpsilon(t.W, other.W)
}

// Add returns a new Tuple which results from adding the other Tuple to the
// first
func Add(t, other Tuple) Tuple {
	return NewTuple(t.X+other.X,
		t.Y+other.Y,
		t.Z+other.Z,
		t.W+other.W)
}

// Subtract returns a new Tuple which results from subtracting the other Tuple to the
// first
func Subtract(t, other Tuple) Tuple {
	return NewTuple(t.X-other.X,
		t.Y-other.Y,
		t.Z-other.Z,
		t.W-other.W)
}

// Negate returns a new Tuple which negates each component of the Tuple
func Negate(t Tuple) Tuple {
	return NewTuple(-t.X, -t.Y, -t.Z, -t.W)
}

// Multiply returns a new Tuple which represents the Tuple multiplied by a
// scalar value
func Multiply(t Tuple, n float64) Tuple {
	return NewTuple(t.X*n, t.Y*n, t.Z*n, t.W*n)
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
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Dot or scalar product returns a scalar value from two Vectors, can be used
// to represent the angle between two Vectors
func Dot(t, other Tuple) float64 {
	return (t.X * other.X) +
		(t.Y * other.Y) +
		(t.Z * other.Z) +
		(t.W * other.W)
}

// Cross or Vector product returns a new Vector that is perpendicular to both
// Vectors
func Cross(t, other Tuple) Tuple {
	return NewVector((t.Y*other.Z)-(t.Z*other.Y),
		(t.Z*other.X)-(t.X*other.Z),
		(t.X*other.Y)-(t.Y*other.X))
}

const epsilon = 0.0001

func inEpsilon(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
