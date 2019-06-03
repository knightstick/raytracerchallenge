package tuples_test

import (
	"math"
	"testing"

	"github.com/knightstick/raytracerchallenge/tuples"
)

const epsilon = 0.00001

func TestTuples(t *testing.T) {
	t.Run("A Tuple with w=1.0 is a point", func(t *testing.T) {
		a := tuples.Point(tuples.New(4.3, -4.2, 3.1, 1.0))

		assertPointEquals(a, 4.3, -4.2, 3.1, 1.0, t)

		if !tuples.IsPoint(a) {
			t.Errorf("expected %v to be a Point", a)
		}

		if tuples.IsVector(a) {
			t.Errorf("expected %v to be a Point", a)
		}
	})

	t.Run("A tuple with w=0 is a Vector", func(t *testing.T) {
		a := tuples.Vector(tuples.New(4.3, -4.2, 3.1, 0.0))

		assertPointEquals(a, 4.3, -4.2, 3.1, 0.0, t)

		if tuples.IsPoint(a) {
			t.Errorf("expected %v not to be a Point", a)
		}

		if !tuples.IsVector(a) {
			t.Errorf("expected %v to be a Vector", a)
		}
	})

	t.Run("NewPoint creates tuples with w=1", func(t *testing.T) {
		point := tuples.NewPoint(4, -4, 3)

		assertPointEquals(point, 4, -4, 3, 1, t)

		if !tuples.IsPoint(point) {
			t.Errorf("expected %v to be a Point", point)
		}
	})

	t.Run("NewVector creates tuples with w=0", func(t *testing.T) {
		point := tuples.NewVector(4.3, -4.2, 3.1)

		assertPointEquals(point, 4.3, -4.2, 3.1, 0, t)

		if !tuples.IsVector(point) {
			t.Errorf("expected %v to be a Vector", point)
		}
	})
}

func TestTupleEquality(t *testing.T) {
	t.Run("Two Tuples are equal if all the values are very close", func(t *testing.T) {
		t1 := tuples.New(1.0, 2.5, 4.999, 0)
		t2 := tuples.New(1.0000001, 2.50000056, 4.999000008, 0)

		if !tuples.Equal(t1, t2) {
			t.Errorf("expected %v to be equal to %v", t1, t2)
		}
	})

	t.Run("Two Tuples are not equal if the values are different", func(t *testing.T) {
		t1 := tuples.New(1.0, 2.5, 4.999, 0)
		t2 := tuples.New(5.0, 2.5, 4.999, 0)

		if tuples.Equal(t1, t2) {
			t.Errorf("expected %v not to be equal to %v", t1, t2)
		}
	})
}

func TestAddition(t *testing.T) {
	a1 := tuples.New(3, -2, 5, 1)
	a2 := tuples.New(-2, 3, 1, 0)

	sum := tuples.Add(a1, a2)
	expected := tuples.New(1, 1, 6, 1)

	if !tuples.Equal(sum, expected) {
		t.Errorf("expected %v, but got %v", expected, sum)
	}
}

func TestSubtraction(t *testing.T) {
	t.Run("Subtracting two points", func(t *testing.T) {
		p1 := tuples.NewPoint(3, 2, 1)
		p2 := tuples.NewPoint(5, 6, 7)

		difference := tuples.Subtract(p1, p2)
		expected := tuples.NewVector(-2, -4, -6)

		assertEqual(difference, expected, t)
	})

	t.Run("Subtracting a Vector from a point", func(t *testing.T) {
		p := tuples.NewPoint(3, 2, 1)
		v := tuples.NewVector(5, 6, 7)

		difference := tuples.Subtract(p, v)
		expected := tuples.NewPoint(-2, -4, -6)

		assertEqual(difference, expected, t)
	})

	t.Run("Subtracting two Vectors", func(t *testing.T) {
		v1 := tuples.NewVector(3, 2, 1)
		v2 := tuples.NewVector(5, 6, 7)

		difference := tuples.Subtract(v1, v2)
		expected := tuples.NewVector(-2, -4, -6)

		assertEqual(difference, expected, t)
	})
}

func TestNegation(t *testing.T) {
	t.Run("Subtracting a Vector from the zero Vector", func(t *testing.T) {
		zero := tuples.NewVector(0, 0, 0)
		v := tuples.NewVector(1, -2, 3)

		assertEqual(tuples.Subtract(zero, v), tuples.NewVector(-1, 2, -3), t)
	})

	t.Run("Negating a Tuple", func(t *testing.T) {
		v := tuples.New(1, -2, 3, -4)

		assertEqual(tuples.Negate(v), tuples.New(-1, 2, -3, 4), t)
	})
}

func TestMultiplicationAndDivision(t *testing.T) {
	t.Run("Multiplying a Tuple by a scalar", func(t *testing.T) {
		a := tuples.New(1, -2, 3, -4)

		assertEqual(tuples.Multiply(a, 3.5), tuples.New(3.5, -7, 10.5, -14), t)
	})

	t.Run("Multiplying a Tuple by a fraction", func(t *testing.T) {
		a := tuples.New(1, -2, 3, -4)

		assertEqual(tuples.Multiply(a, 0.5), tuples.New(0.5, -1, 1.5, -2), t)
	})

	t.Run("Dividing a Tuple by a scalar", func(t *testing.T) {
		a := tuples.New(1, -2, 3, -4)

		assertEqual(tuples.Divide(a, 2), tuples.New(0.5, -1, 1.5, -2), t)
	})
}

func TestMagnitude(t *testing.T) {
	t.Run("Computing the magnitude of Vector(1,0,0)", func(t *testing.T) {
		v := tuples.NewVector(1, 0, 0)

		magnitude := tuples.Magnitude(v)

		if magnitude != 1 {
			t.Errorf("expected magnitude of %v to be 1, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(0,1,0)", func(t *testing.T) {
		v := tuples.NewVector(0, 1, 0)

		magnitude := tuples.Magnitude(v)

		if magnitude != 1 {
			t.Errorf("expected magnitude of %v to be 1, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(0,0,1)", func(t *testing.T) {
		v := tuples.NewVector(0, 0, 1)

		magnitude := tuples.Magnitude(v)

		if magnitude != 1 {
			t.Errorf("expected magnitude of %v to be 1, got %f", v, magnitude)
		}
	})

	t.Run("Computing the magnitude of Vector(1,2,3)", func(t *testing.T) {
		v := tuples.NewVector(1, 2, 3)

		magnitude := tuples.Magnitude(v)

		assertInEpsilon(magnitude, math.Sqrt(14), t)
	})

	t.Run("Computing the magnitude of Vector(-1,-2,-3)", func(t *testing.T) {
		v := tuples.NewVector(-1, -2, -3)

		magnitude := tuples.Magnitude(v)

		assertInEpsilon(magnitude, math.Sqrt(14), t)
	})
}

func TestNormalizing(t *testing.T) {
	t.Run("Normalizing Vector (4, 0, 0) gives (1, 0, 0)", func(t *testing.T) {
		v := tuples.NewVector(4, 0, 0)

		assertEqual(tuples.Normalize(v), tuples.NewVector(1, 0, 0), t)
	})

	t.Run("Normalizing Vector (1, 2, 3)", func(t *testing.T) {
		v := tuples.NewVector(1, 2, 3)

		expected := tuples.NewVector(0.26726, 0.53452, 0.80178)

		assertEqual(tuples.Normalize(v), tuples.Tuple(expected), t)
	})

	t.Run("The magnitude of a normalized Vector", func(t *testing.T) {
		v := tuples.NewVector(1, 2, 3)
		norm := tuples.Normalize(v)

		actualMagnitude := tuples.Magnitude(norm)

		if actualMagnitude != 1 {
			t.Errorf("expected %v to have magnitude 1, but got %f", norm, actualMagnitude)
		}
	})
}

func TestDotProduct(t *testing.T) {
	t.Run("The dot product of two tuples", func(t *testing.T) {
		a := tuples.NewVector(1, 2, 3)
		b := tuples.NewVector(2, 3, 4)
		dot := tuples.Dot(a, b)

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

		assertEqual(tuples.Cross(a, b), tuples.NewVector(-1, 2, -1), t)
		assertEqual(tuples.Cross(b, a), tuples.NewVector(1, -2, 1), t)
	})
}

func TestColors(t *testing.T) {
	t.Run("Colors are (red, green, blue) tuples", func(t *testing.T) {
		c := tuples.NewColor(-0.5, 0.4, 1.7)

		assertInEpsilon(c.Red(), -0.5, t)
		assertInEpsilon(c.Green(), 0.4, t)
		assertInEpsilon(c.Blue(), 1.7, t)
	})
}

func TestColorOperations(t *testing.T) {
	t.Run("Adding colors", func(t *testing.T) {
		c1 := tuples.NewColor(0.9, 0.6, 0.75)
		c2 := tuples.NewColor(0.7, 0.1, 0.25)

		sum := tuples.Add(c1, c2)

		assertEqual(sum, tuples.NewColor(1.6, 0.7, 1.0), t)
	})

	t.Run("Subtracting colors", func(t *testing.T) {
		c1 := tuples.NewColor(0.9, 0.6, 0.75)
		c2 := tuples.NewColor(0.7, 0.1, 0.25)

		difference := tuples.Subtract(c1, c2)

		assertEqual(difference, tuples.NewColor(0.2, 0.5, 0.5), t)
	})

	t.Run("Multiplying a color by a scalar", func(t *testing.T) {
		c := tuples.NewColor(0.2, 0.3, 0.4)

		product := tuples.Multiply(c, 2)

		assertEqual(product, tuples.NewColor(0.4, 0.6, 0.8), t)
	})

	t.Run("Multiplying colors", func(t *testing.T) {
		c1 := tuples.NewColor(1, 0.2, 0.4)
		c2 := tuples.NewColor(0.9, 1, 0.1)

		product := tuples.MultiplyColors(c1, c2)

		assertEqual(product, tuples.NewColor(0.9, 0.2, 0.04), t)
	})
}

func assertEqual(actual, expected tuples.Tuplelike, t *testing.T) {
	t.Helper()

	if !tuples.Equal(actual, expected) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func assertPointEquals(actual tuples.Pointlike, x, y, z, w float64, t *testing.T) {
	t.Helper()

	if actual.X() != x {
		t.Errorf("expected %v.X() to equal %f", actual, actual.X())
	}

	if actual.Y() != y {
		t.Errorf("expected %v.X() to equal %f", actual, actual.X())
	}

	if actual.Z() != z {
		t.Errorf("expected %v.X() to equal %f", actual, actual.X())
	}

	if actual.W() != w {
		t.Errorf("expected %v.X() to equal %f", actual, actual.X())
	}
}

func assertInEpsilon(actual, expected float64, t *testing.T) {
	t.Helper()

	if approximatelyEqual(actual, expected) {
		t.Errorf("expected %f to be approximately equal to %f", actual, expected)
	}
}

func approximatelyEqual(a, b float64) bool {
	return math.Abs(a-b) > epsilon
}
