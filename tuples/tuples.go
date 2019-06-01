package tuples

// A Tuple is a list of 4 floats used to represent Points and Vectors
type Tuple struct {
	X float32
	Y float32
	Z float32
	W float32
}

// NewTuple instantiates a new Tuple
func NewTuple(x, y, z, w float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

// IsPoint checks if a Tuple is a Point
func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector checks if a Tuple is a Vector
func (t Tuple) IsVector() bool {
	return t.W == 0
}
