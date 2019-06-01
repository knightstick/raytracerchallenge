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

	t.Run("A tuple with w=0 is a vector", func(t *testing.T) {
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
