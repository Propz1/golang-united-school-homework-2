package square

import (
	"math"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	var x float64 = 10
	var result float64

	t.Run("Triangle", func(t *testing.T) {

		result = math.Sqrt(3) * 25

		realResult := CalcSquare(x, SidesTriangle)

		if result != realResult {
			t.Errorf("expected result %f != %f real result", realResult, result)
		}
	})

	t.Run("Square", func(t *testing.T) {

		result = 100

		realResult := CalcSquare(x, SidesSquare)

		if result != realResult {
			t.Errorf("expected result %f != %f real result", realResult, result)
		}
	})

	t.Run("Circle", func(t *testing.T) {

		result = Pi * 100

		realResult := CalcSquare(x, SidesCircle)

		if result != realResult {
			t.Errorf("expected result %f != %f real result", realResult, result)
		}
	})

}
