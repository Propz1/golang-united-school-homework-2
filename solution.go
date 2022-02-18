package square

import (
	"math"
)

type intCustomType int

const (
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
	SidesCircle   intCustomType = 0
	Pi            float64       = math.Pi
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {

	if sideLen < 0 {
		return 0
	}

	switch sidesNum {
	case 0:
		return Pi * math.Pow(sideLen, 2)
	case 3:
		return math.Sqrt(3) * math.Pow(sideLen, 2) / 4
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
