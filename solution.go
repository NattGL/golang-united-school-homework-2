package square

import (
	"math"
)

type SidesNumber int

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * sideLen * sideLen
	case 3:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
