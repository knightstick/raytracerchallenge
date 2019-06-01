package tuples_test

import (
	"testing"

	"github.com/knightstick/raytracerchallenge/tuples"
)

func TestTuples(t *testing.T) {
	t.Run("A Tuple can be a Point", func(t *testing.T) {
		a := tuples.NewTuple(4.3, -4.2, 3.1, 1.0)

		assertEqual(a.X, 4.3, t)
		assertEqual(a.Y, -4.2, t)
		assertEqual(a.Z, 3.1, t)
		assertEqual(a.W, 1.0, t)

		if !a.IsPoint() {
			t.Errorf("expected %v to be a Point", a)
		}

		if a.IsVector() {
			t.Errorf("expected %v not to be a Vector", a)
		}
	})

	t.Run("A Tuple can be a Point", func(t *testing.T) {
		a := tuples.NewTuple(4.3, -4.2, 3.1, 0.0)

		assertEqual(a.X, 4.3, t)
		assertEqual(a.Y, -4.2, t)
		assertEqual(a.Z, 3.1, t)
		assertEqual(a.W, 0.0, t)

		if a.IsPoint() {
			t.Errorf("expected %v not to be a Point", a)
		}

		if !a.IsVector() {
			t.Errorf("expected %v to be a Vector", a)
		}
	})
}

func assertEqual(actual, expected float32, t *testing.T) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %b to equal %b", actual, expected)
	}
}
