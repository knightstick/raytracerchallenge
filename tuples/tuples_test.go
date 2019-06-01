package tuples_test

import (
	"math"
	"testing"

	"github.com/knightstick/raytracerchallenge/tuples"
)

const epsilon = 0.00001

func TestTuples(t *testing.T) {
	t.Run("A Tuple with w=1.0 is a point", func(t *testing.T) {
		a := tuples.NewTuple(4.3, -4.2, 3.1, 1.0)

		assertTupleEqual(a, 4.3, -4.2, 3.1, 1.0, t)

		if !a.IsPoint() {
			t.Errorf("expected %v to be a Point", a)
		}

		if a.IsVector() {
			t.Errorf("expected %v not to be a Vector", a)
		}
	})

	t.Run("A tuple with w=0 is a Vector", func(t *testing.T) {
		a := tuples.NewTuple(4.3, -4.2, 3.1, 0.0)

		assertTupleEqual(a, 4.3, -4.2, 3.1, 0.0, t)

		if a.IsPoint() {
			t.Errorf("expected %v not to be a Point", a)
		}

		if !a.IsVector() {
			t.Errorf("expected %v to be a Vector", a)
		}
	})

	t.Run("NewPoint creates tuples with w=1", func(t *testing.T) {
		point := tuples.NewPoint(4, -4, 3)

		assertTupleEqual(point, 4, -4, 3, 1, t)
	})

	t.Run("NewVector creates tuples with w=0", func(t *testing.T) {
		point := tuples.NewVector(4.3, -4.2, 3.1)

		if !point.IsVector() {
			t.Errorf("expected %v to be a Vector", point)
		}
	})
}

func TestTupleEquality(t *testing.T) {
	t.Run("Two Tuples are equal if all the values are very close", func(t *testing.T) {
		t1 := tuples.NewTuple(1.0, 2.5, 4.999, 0)
		t2 := tuples.NewTuple(1.0000001, 2.50000056, 4.999000008, 0)

		if !t1.Equal(t2) {
			t.Errorf("expected %v to be equal to %v", t1, t2)
		}
	})

	t.Run("Two Tuples are not equal if the values are different", func(t *testing.T) {
		t1 := tuples.NewTuple(1.0, 2.5, 4.999, 0)
		t2 := tuples.NewTuple(5.0, 2.5, 4.999, 0)

		if t1.Equal(t2) {
			t.Errorf("expected %v not to be equal to %v", t1, t2)
		}
	})
}

func TestAddition(t *testing.T) {
	a1 := tuples.NewTuple(3, -2, 5, 1)
	a2 := tuples.NewTuple(-2, 3, 1, 0)

	sum := a1.Add(a2)
	expected := tuples.NewTuple(1, 1, 6, 1)

	if !sum.Equal(expected) {
		t.Errorf("expected %v, but got %v", expected, sum)
	}
}

func TestSubtraction(t *testing.T) {
	t.Run("Subtracting two points", func(t *testing.T) {
		p1 := tuples.NewPoint(3, 2, 1)
		p2 := tuples.NewPoint(5, 6, 7)

		difference := p1.Subtract(p2)
		expected := tuples.NewVector(-2, -4, -6)

		assertEqual(difference, expected, t)
	})

	t.Run("Subtracting a Vector from a point", func(t *testing.T) {
		p := tuples.NewPoint(3, 2, 1)
		v := tuples.NewVector(5, 6, 7)

		difference := p.Subtract(v)
		expected := tuples.NewPoint(-2, -4, -6)

		assertEqual(difference, expected, t)
	})

	t.Run("Subtracting two Vectors", func(t *testing.T) {
		v1 := tuples.NewVector(3, 2, 1)
		v2 := tuples.NewVector(5, 6, 7)

		difference := v1.Subtract(v2)
		expected := tuples.NewVector(-2, -4, -6)

		assertEqual(difference, expected, t)
	})
}

func TestNegation(t *testing.T) {
	t.Run("Subtracting a Vector from the zero Vector", func(t *testing.T) {
		zero := tuples.NewVector(0, 0, 0)
		v := tuples.NewVector(1, -2, 3)

		assertEqual(zero.Subtract(v), tuples.NewVector(-1, 2, -3), t)
	})

	t.Run("Negating a Tuple", func(t *testing.T) {
		v := tuples.NewTuple(1, -2, 3, -4)

		assertEqual(v.Negate(), tuples.NewTuple(-1, 2, -3, 4), t)
	})
}

func TestMultiplicationAndDivision(t *testing.T) {
	t.Run("Multiplying a Tuple by a scalar", func(t *testing.T) {
		a := tuples.NewTuple(1, -2, 3, -4)

		assertEqual(a.Multiply(3.5), tuples.NewTuple(3.5, -7, 10.5, -14), t)
	})

	t.Run("Multiplying a Tuple by a fraction", func(t *testing.T) {
		a := tuples.NewTuple(1, -2, 3, -4)

		assertEqual(a.Multiply(0.5), tuples.NewTuple(0.5, -1, 1.5, -2), t)
	})

	t.Run("Dividing a Tuple by a scalar", func(t *testing.T) {
		a := tuples.NewTuple(1, -2, 3, -4)

		assertEqual(a.Divide(2), tuples.NewTuple(0.5, -1, 1.5, -2), t)
	})
}

func TestMagnitude(t *testing.T) {
	t.Run("Computing the magnitude of Vector(1,0,0)", func(t *testing.T) {
		v := tuples.NewVector(1, 0, 0)

		magnitude := v.Magnitude()

		if magnitude != 1 {
			t.Errorf("expected magnitude of %v to be 1, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(0,1,0)", func(t *testing.T) {
		v := tuples.NewVector(0, 1, 0)

		magnitude := v.Magnitude()

		if magnitude != 1 {
			t.Errorf("expected magnitude of %v to be 1, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(0,0,1)", func(t *testing.T) {
		v := tuples.NewVector(0, 0, 1)

		magnitude := v.Magnitude()

		if magnitude != 1 {
			t.Errorf("expected magnitude of %v to be 1, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(1,2,3)", func(t *testing.T) {
		v := tuples.NewVector(1, 2, 3)

		magnitude := v.Magnitude()

		if magnitude != math.Sqrt(14) {
			t.Errorf("expected magnitude of %v to be √14, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(-1,-2,-3)", func(t *testing.T) {
		v := tuples.NewVector(-1, -2, -3)

		magnitude := v.Magnitude()

		if magnitude != math.Sqrt(14) {
			t.Errorf("expected magnitude of %v to be √14, got %f", v, magnitude)
		}
	})
}

func TestNormalizing(t *testing.T) {
	t.Run("Normalizing Vector (4, 0, 0) gives (1, 0, 0)", func(t *testing.T) {
		v := tuples.NewVector(4, 0, 0)

		assertEqual(v.Normalize(), tuples.NewVector(1, 0, 0), t)
	})

	t.Run("Normalizing Vector (1, 2, 3)", func(t *testing.T) {
		v := tuples.NewVector(1, 2, 3)

		expected := tuples.NewVector(0.26726, 0.53452, 0.80178)

		assertEqual(v.Normalize(), expected, t)
	})

	t.Run("The magnitude of a normalized Vector", func(t *testing.T) {
		v := tuples.NewVector(1, 2, 3)
		norm := v.Normalize()

		actualMagnitude := norm.Magnitude()

		if norm.Magnitude() != 1 {
			t.Errorf("expected %v to have magnitude 1, but got %f", norm, actualMagnitude)
		}
	})
}

func TestDotProduct(t *testing.T) {
	t.Run("The dot product of two tuples", func(t *testing.T) {
		a := tuples.NewVector(1, 2, 3)
		b := tuples.NewVector(2, 3, 4)
		dot := a.Dot(b)

		if dot != 20 {
			t.Errorf("expected dot product of %v and %v to equal 20, but got %f",
				a, b, dot)
		}
	})
}

func TestCrossProduct(t *testing.T) {
	t.Run("The cross product of two Vectors", func(t *testing.T) {
		a := tuples.NewVector(1, 2, 3)
		b := tuples.NewVector(2, 3, 4)

		assertEqual(a.Cross(b), tuples.NewVector(-1, 2, -1), t)
		assertEqual(b.Cross(a), tuples.NewVector(1, -2, 1), t)
	})
}

func assertEqual(actual, expected tuples.Tuple, t *testing.T) {
	t.Helper()

	if !actual.Equal(expected) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func assertInEpsilon(actual, expected float64, t *testing.T) {
	t.Helper()

	if approximatelyEqual(actual, expected) {
		t.Errorf("expected %f to be approximately equal to %f", actual, expected)
	}
}

func assertTupleEqual(tup tuples.Tuple, x, y, z, w float64, t *testing.T) {
	t.Helper()

	assertInEpsilon(tup.X, x, t)
	assertInEpsilon(tup.Y, y, t)
	assertInEpsilon(tup.Z, z, t)
	assertInEpsilon(tup.W, w, t)
}

func approximatelyEqual(a, b float64) bool {
	return math.Abs(a-b) > epsilon
}
