package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SideNumber int

const (
	SidesTriangle SideNumber = 3
	SidesSquare   SideNumber = 4
	SidesCircle   SideNumber = 0
)

func CalcSquare(sideLen float64, sidesNum SideNumber) float64 {
	var square float64

	switch sidesNum {
	case SidesTriangle:
		square = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case SidesSquare:
		square = sideLen * sideLen
	case SidesCircle:
		square = math.Pow(sideLen, 2) * math.Pi
	default:
		square = 0
	}
	return square
}
