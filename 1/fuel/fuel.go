package fuel

import (
	"math"
)

func Calculator(mass float64) (fuel int) {
	fuel = int(math.Floor(mass/3) - 2)
	return
}
