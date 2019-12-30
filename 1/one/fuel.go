package main

import (
	"math"
)

// Sum receives a numeric array and calculates each element into a total sum.
func Sum(input []float64) (sum int) {
	for _, v := range input {
		sum += calculator(v)
	}
	return sum
}

func calculator(mass float64) (fuel int) {
	fuel = int(math.Floor(mass/3) - 2)
	return
}
