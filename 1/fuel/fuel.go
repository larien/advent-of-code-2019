package fuel

import (
	"fmt"
	"math"
)

func Calculator(mass float64) int {
	mass /= 3

	mass = math.Floor(mass)

	mass -= 2

	fmt.Println(mass)
	return int(mass)
}
