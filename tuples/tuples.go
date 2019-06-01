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
func (t Tuple) Equal(other Tuple) bool {
	return inEpsilon(t.X, other.X) && inEpsilon(t.Y, other.Y) &&
		inEpsilon(t.Z, other.Z) && inEpsilon(t.W, other.W)
}

// Add returns a new Tuple which results from adding the other Tuple to the
// first
func (t Tuple) Add(other Tuple) Tuple {
	return NewTuple(t.X+other.X,
		t.Y+other.Y,
		t.Z+other.Z,
		t.W+other.W)
}

// Subtract returns a new Tuple which results from subtracting the other Tuple to the
// first
func (t Tuple) Subtract(other Tuple) Tuple {
	return NewTuple(t.X-other.X,
		t.Y-other.Y,
		t.Z-other.Z,
		t.W-other.W)
}

// Negate returns a new Tuple which negates each component of the Tuple
func (t Tuple) Negate() Tuple {
	return NewTuple(-t.X, -t.Y, -t.Z, -t.W)
}

const epsilon = 0.0001

func inEpsilon(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
