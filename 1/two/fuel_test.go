package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		inputs   []float64
		expected float64
	}{
		{[]float64{14, 0}, 2},
		{[]float64{14, 1969}, 968},
		{[]float64{14, 1969, 100756}, 51314},
	}
	for index, test := range cases {
		testName := fmt.Sprintf("Sum (%d)", index)
		t.Run(testName, func(t *testing.T) {
			result := Sum(test.inputs)
			assert(t, result, test.expected)
		})
	}
}

func TestCalculator(t *testing.T) {
	cases := []struct {
		mass     float64
		fuel     float64
		expected float64
	}{
		{14, 0, 2},
		{1969, 0, 966},
		{100756, 0, 50346},
	}
	for index, test := range cases {
		testName := fmt.Sprintf("Calculator (%d)", index)
		t.Run(testName, func(t *testing.T) {
			result := calculator(test.mass, test.fuel)
			assert(t, result, test.expected)
		})
	}
}

func assert(t *testing.T, result, expected float64) {
	t.Helper()
	if result != expected {
		t.Errorf("result: %v, expected: %v", result, expected)
	}
}
