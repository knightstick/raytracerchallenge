package tuples_test

import (
	"testing"

	"github.com/knightstick/raytracerchallenge/tuples"
)

func TestTuples(t *testing.T) {
	t.Run("A Tuple can be a Point", func(t *testing.T) {
		a := tuples.NewTuple(4.3, -4.2, 3.1, 1.0)

		assertTupleEqual(a, 4.3, -4.2, 3.1, 1.0, t)

		if !a.IsPoint() {
			t.Errorf("expected %v to be a Point", a)
		}

		if a.IsVector() {
			t.Errorf("expected %v not to be a Vector", a)
		}
	})

	t.Run("A Tuple can be a Point", func(t *testing.T) {
		a := tuples.NewTuple(4.3, -4.2, 3.1, 0.0)

		assertTupleEqual(a, 4.3, -4.2, 3.1, 0.0, t)

		if a.IsPoint() {
			t.Errorf("expected %v not to be a Point", a)
		}

		if !a.IsVector() {
			t.Errorf("expected %v to be a Vector", a)
		}
	})

	t.Run("NewPoint creates a Point", func(t *testing.T) {
		point := tuples.NewPoint(4.3, -4.2, 3.1)

		if !point.IsPoint() {
			t.Errorf("expected %v to be a Point", point)
		}
	})

	t.Run("NewVector creates a Vector", func(t *testing.T) {
		point := tuples.NewVector(4.3, -4.2, 3.1)

		if !point.IsVector() {
			t.Errorf("expected %v to be a Vector", point)
		}
	})
}

func assertEqual(actual, expected float32, t *testing.T) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %b to equal %b", actual, expected)
	}
}

func assertTupleEqual(tup tuples.Tuple, x, y, z, w float64, t *testing.T) {
	t.Helper()

	assertEqual(tup.X, 4.3, t)
	assertEqual(tup.Y, -4.2, t)
	assertEqual(tup.Z, 3.1, t)
	assertEqual(tup.W, 1.0, t)
}