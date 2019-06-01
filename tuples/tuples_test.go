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
