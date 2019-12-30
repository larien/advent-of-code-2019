package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		inputs   []float64
		expected int
	}{
		{[]float64{12, 14}, 4},
		{[]float64{12, 14, 1969}, 658},
		{[]float64{12, 14, 1969, 100756}, 34241},
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
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for index, test := range cases {
		testName := fmt.Sprintf("Calculator (%d)", index)
		t.Run(testName, func(t *testing.T) {
			result := calculator(test.mass)
			assert(t, result, test.expected)
		})
	}
}

func assert(t *testing.T, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("result: %v, expected %v", result, expected)
	}
}
