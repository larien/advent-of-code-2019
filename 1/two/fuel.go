package main

import (
	"math"
)

// Sum receives a numeric array and calculates each element into a total sum.
func Sum(input []float64) (sum float64) {
	for _, v := range input {
		var fuel float64
		fuel += calculator(v, fuel)
		sum += fuel
	}
	return sum
}

func calculator(mass, fuel float64) float64 {
	calculation := math.Floor(mass/3) - 2
	if calculation > 0 {
		fuel += calculation
		return calculator(calculation, fuel)
	}
	return fuel
}
